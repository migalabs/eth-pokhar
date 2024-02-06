# eth-pokhar

Tool to identify validators &lt;> entities in the Ethereum consensus layer

## Database migrations

In case you encounter any issue with the database, you can force the database version using the golang-migrate command line. Please refer [here](https://github.com/golang-migrate/migrate) for more information.
More specifically, one could clean the migrations by forcing the version with <br>
`migrate -path / -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force <current_version>` <br>
If specific upgrades or downgrades need to be done manually, one could do this with <br>
`migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up`
