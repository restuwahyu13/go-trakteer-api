package configs

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func Connection(driver string) *sqlx.DB {

	db, _ := sqlx.Open(driver, viper.GetString("PG_DSN"))
	return db
}
