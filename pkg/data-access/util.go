package dataaccess

import (
	"database/sql"
	"student-service/pkg/config"
	"student-service/pkg/data-access/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func InitializeSequelDB() *bun.DB {
	// initialize pgx config
	pgxConfig, err := pgx.ParseConfig(config.AppConfig.DBUrl)
	if err != nil {
		panic(err)
	}

	// use simple protocol as bun does not benefit from using implicit prepared statements
	pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	var sqlDB *sql.DB = stdlib.OpenDB(*pgxConfig)

	db := bun.NewDB(sqlDB, pgdialect.New())
	db.RegisterModel((*dto.StudentClass)(nil))

	return db
}
