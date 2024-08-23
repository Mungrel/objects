package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/pkg/errors"
	"github.com/snikch/goose/lib/goose"
)

const (
	driver = "mysql"
	dsn    = "app:password@tcp(localhost:3306)/app?parseTime=true&clientFoundRows=true"
)

var cnf = &goose.DBConf{
	MigrationsDir: "./db/migrations",
	Driver: goose.DBDriver{
		Name:    driver,
		OpenStr: dsn,
		Dialect: goose.MySqlDialect{},
	},
}

func Client() *sqlx.DB {
	client := sqlx.MustConnect(driver, dsn)
	client.Mapper = reflectx.NewMapper("db")
	return client
}

func Migrate() error {
	version, err := goose.GetMostRecentDBVersion(cnf.MigrationsDir)
	if err != nil {
		return errors.Wrap(err, "failed to get most recent DB version")
	}

	if err = goose.RunMigrations(cnf, "./db/migrations", version); err != nil {
		return errors.Wrap(err, "failed to run migrations")
	}

	return nil
}
