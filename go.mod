module database-manager-cli

go 1.18

require (
	github.com/DATA-DOG/go-sql-mock v1.5.0
	github.com/docker/docker v20.10.7+incompatible
	github.com/google/uuid v1.3.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.10.3
	github.com/prisma/prisma-client-go v1.14.0
	github.com/redis/go-redis/v9 v9.1.0
	gopkg.in/yaml.v3 v3.0.0-20210107162922-496545a6307b
)

replace github.com/docker/docker => github.com/docker/docker v20.10.7-0.20210616153541-5ef8cf75a5b7
