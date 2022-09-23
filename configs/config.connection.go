package configs

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func Connection(driver string) *sqlx.DB {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	db, _ := sqlx.ConnectContext(ctx, driver, viper.GetString("PG_DSN"))
	return db
}
