package db

import "github.com/pkg/errors"

const (
	identifyCoinbaseValidatorsQuery = `
	UPDATE t_identified_validators
	SET f_pool_name = 'coinbase'
	WHERE f_validator_pubkey in (
		SELECT distinct f_validator_pubkey from t_beacon_depositors_transactions
		JOIN t_beacon_deposits on t_beacon_deposits.f_tx_hash=t_beacon_depositors_transactions.f_tx_hash
		WHERE 
		f_to = '00000000219ab540356cbb839cbe05303d7705fa'
		AND f_from in (
			select f_from from t_beacon_depositors_transactions 
			where 
			f_from in (
				SELECT 
				f_from
				FROM 
				t_beacon_depositors_transactions 
				where 
				f_to = '00000000219ab540356cbb839cbe05303d7705fa' 
				group by 
				f_from 
				having 
				count(*) = 1
			) 
			AND f_to in(
				'a090e606e30bd747d4e6245a1517ebe430f0057e',
				'71660c4005ba85c37ccec55d0c4493e66fe775d3',
				'a9d1e08c7793af67e9d92fe308d5697fb81d3e43',
				'77696bb39917c91a0c3908d577d5e322095425ca',
				'7c195d981abfdc3ddecd2ca0fed0958430488e34',
				'95a9bd206ae52c4ba8eecfc93d18eacdd41c88cc',
				'b739d0895772dbb71a89a3754a160269068f0d45',
				'503828976d22510aad0201ac7ec88293211d23da',
				'ddfabcdc4d8ffc6d5beaf154f18b778f892a0740',
				'3cd751e6b0078be393132286c442345e5dc49699',
				'b5d85cbf7cb3ee0d56b3bb207d5fc4b82f43f511',
				'eb2629a2734e272bcc07bda959863f316f4bd4cf',
				'd688aea8f7d450909ade10c47faa95707b0682d9',
				'02466e547bfdab679fc49e96bbfc62b9747d997c',
				'6b76f8b1e9e59913bfe758821887311ba1805cab',
				'be3c68821d585cf1552214897a1c091014b1eb0a',
				'f6874c88757721a02f47592140905c4336dfbc61',
				'881d4032abe4188e2237efcd27ab435e81fc6bb1',
				'6c8dd0e9cc58c07429e065178d88444b60e60b80',
				'bc8ec259e3026ae0d87bc442d034d6882ce4a35c',
				'02d24cab4f2c3bf6e6eb07ea07e45f96baccffe7',
				'ce352e98934499be70f641353f16a47d9e1e3abd',
				'90e18a6920985dbacc3d76cf27a3f2131923c720',
				'4b23d52eff7c67f5992c2ab6d3f69b13a6a33561',
				'd2276af80582cac230edc4c42e9a9c096f3c09aa')
		)
	)
	`
)

func (p *PostgresDBService) IdentifyCoinbaseValidators() error {
	_, err := p.psqlPool.Query(p.ctx, identifyCoinbaseValidatorsQuery)
	if err != nil {
		return errors.Wrap(err, "error identifying coinbase validators")
	}
	return nil
}
