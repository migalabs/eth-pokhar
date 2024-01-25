package db

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/sirupsen/logrus"
)

// Static postgres queries, for each modification in the tables, the table needs to be reseted
var (
	// wlogrus associated with the postgres db
	modName  = "db"
	PsqlType = "postgres-db"
	wlog     = logrus.WithField(
		"module", PsqlType,
	)
	MAX_BATCH_QUEUE       = 1000
	MAX_EPOCH_BATCH_QUEUE = 1
	RoutineFlushTimeout   = time.Duration(1 * time.Second)
)

type PostgresDBServiceOption func(*PostgresDBService) error

type PostgresDBService struct {
	// Control Variables
	ctx           context.Context
	connectionUrl string // the url might not be necessary (better to remove it?¿)
	psqlPool      *pgxpool.Pool
	wgDBWriters   sync.WaitGroup
	writerBatches []*QueryBatch

	writeChan chan Model // Receive tasks to persist
	stop      bool
	workerNum int
}

func New(ctx context.Context, url string, options ...PostgresDBServiceOption) (*PostgresDBService, error) {
	var err error
	pService := &PostgresDBService{
		ctx:           ctx,
		connectionUrl: url,
		workerNum:     1,
	}

	for _, o := range options {
		err := o(pService)
		if err != nil {
			return pService, err
		}
	}

	pService.writerBatches = make([]*QueryBatch, pService.workerNum)
	pService.writeChan = make(chan Model, pService.workerNum)
	return pService, err
}

func WithUrl(url string) PostgresDBServiceOption {
	return func(s *PostgresDBService) error {
		s.connectionUrl = url
		return nil
	}
}

func WithWorkers(workerNum int) PostgresDBServiceOption {

	return func(s *PostgresDBService) error {
		s.workerNum = workerNum
		if s.workerNum < 1 {
			return fmt.Errorf("cannot set a negative number of workers")
		}
		return nil
	}
}

// Connect to the PostgreSQL Database and get the multithread-proof connection
// from the given url-composed credentials
func (s *PostgresDBService) Connect() {
	// spliting the url to don't share any confidential information on wlogs
	wlog.Infof("Connecting to postgres DB %s", s.connectionUrl)
	if strings.Contains(s.connectionUrl, "@") {
		wlog.Debugf("Connecting to PostgresDB at %s", strings.Split(s.connectionUrl, "@")[1])
	}
	psqlPool, err := pgxpool.New(s.ctx, s.connectionUrl)
	if err != nil {
		wlog.Fatalf("could not connect to database: %s", err.Error())
	}
	s.psqlPool = psqlPool
	if strings.Contains(s.connectionUrl, "@") {
		wlog.Infof("PostgresDB %s succesfully connected", strings.Split(s.connectionUrl, "@")[1])
	}

	// init the psql db
	s.makeMigrations()
	go s.runWriters()
}

func (p *PostgresDBService) Finish() {
	p.stop = true
	p.wgDBWriters.Wait()
	wlog.Infof("Routines finished...")
	wlog.Infof("closing connection to database server...")
	p.psqlPool.Close()
	wlog.Infof("connection to database server closed...")
	close(p.writeChan)
}

func (p *PostgresDBService) runWriters() {

	wlog.Info("Launching Beacon State Writers")
	wlog.Infof("Launching %d Beacon State Writers", p.workerNum)
	for i := 0; i < p.workerNum; i++ {
		p.wgDBWriters.Add(1)
		go func(dbWriterID int) {
			defer p.wgDBWriters.Done()
			batcher := NewQueryBatch(p.ctx, p.psqlPool, MAX_BATCH_QUEUE)
			p.writerBatches[dbWriterID] = batcher
			wlogWriter := wlog.WithField("DBWriter", dbWriterID)
			ticker := time.NewTicker(RoutineFlushTimeout)
		loop:
			for {
				select {
				case task := <-p.writeChan:

					var err error
					persis := NewPersistable()

					switch task.Type() {
					case models.TransactionModel:
						q, args := TransactionOperation(task.(models.Transaction))
						persis.query = q
						persis.values = append(persis.values, args...)

					default:
						err = fmt.Errorf("could not figure out the type of write task")
						wlog.Errorf("could not process incoming task, %s", err)
					}
					// ckeck if there is any new query to add
					if !persis.isEmpty() {
						batcher.AddQuery(persis)
					}
					// check if we can flush the batch of queries
					if batcher.IsReadyToPersist() {
						err := batcher.PersistBatch()
						if err != nil {
							wlogWriter.Error("Error processing batch", err.Error())
						}
					}

				case <-p.ctx.Done():
					break loop

				case <-ticker.C:
					// if limit reached or no more queue and pending tasks
					if batcher.IsReadyToPersist() || (len(p.writeChan) == 0 && batcher.Len() > 0) {
						wlog.Tracef("flushing batcher")
						err := batcher.PersistBatch()
						if err != nil {
							wlogWriter.Error("Error processing batch", err.Error())
						}
					}

					if p.stop && len(p.writeChan) == 0 {
						break loop
					}
				}
			}
		}(i)
	}

}

func (p *PostgresDBService) Persist(w Model) {
	p.writeChan <- w
}

type Model interface { // simply to enforce a Model interface
	// For now we simply support insert operations
	Type() models.ModelType // whether insert is activated for this model
}

func (p *PostgresDBService) SingleQuery(query string, args ...interface{}) error {
	rows, err := p.psqlPool.Query(p.ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing query %s: %s", query, err.Error())
	}
	rows.Close()
	return nil
}

func (p *PostgresDBService) GetBatcherStats() []string {
	result := make([]string, 0)

	for idx, batcher := range p.writerBatches {
		result = append(result, fmt.Sprintf("batcher %d=%dms(%d queries)",
			idx,
			batcher.metrics.PersistTime.Milliseconds(),
			batcher.metrics.NumQueries))
	}
	return result
}
