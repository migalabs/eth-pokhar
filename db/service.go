package db

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Static postgres queries, for each modification in the tables, the table needs to be reseted
var (
	// wlogrus associated with the postgres db
	// modName  = "db"
	PsqlType = "postgres-db"
	wlog     = logrus.WithField(
		"module", PsqlType,
	)
)

type PostgresDBServiceOption func(*PostgresDBService) error

type PostgresDBService struct {
	// Control Variables
	ctx             context.Context
	connectionUrl   string // the url might not be necessary (better to remove it?¿)
	psqlPool        *pgxpool.Pool
	writerThreadsWG *sync.WaitGroup
}

func New(ctx context.Context, url string, options ...PostgresDBServiceOption) (*PostgresDBService, error) {
	var err error
	pService := &PostgresDBService{
		ctx:             ctx,
		connectionUrl:   url,
		writerThreadsWG: &sync.WaitGroup{},
	}

	for _, o := range options {
		err := o(pService)
		if err != nil {
			return pService, err
		}
	}

	return pService, err
}

func WithUrl(url string) PostgresDBServiceOption {
	return func(s *PostgresDBService) error {
		s.connectionUrl = url
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
}

func (p *PostgresDBService) Finish() {
	wlog.Debugf("Losing connection to database server...")
	wlog.Debugf("Waiting for all writer threads to finish...")
	p.writerThreadsWG.Wait()
	p.psqlPool.Close()
	wlog.Debugf("Connection to database server closed...")
}

func (p *PostgresDBService) SingleQuery(query string, args ...interface{}) error {
	rows, err := p.psqlPool.Query(p.ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing query %s: %s", query, err.Error())
	}
	rows.Close()
	return nil
}
