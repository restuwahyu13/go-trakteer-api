package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/restuwahyu13/go-trakteer-api/configs"
	"github.com/restuwahyu13/go-trakteer-api/packages"
	"github.com/restuwahyu13/go-trakteer-api/routes"
)

func main() {
	err := packages.Viper()
	if err != nil {
		log.Fatalf(".env file not load: %v", err)
	}

	db := SetupDatabase()
	router := SetupRouter()

	routes.NewUsersRoute("/api/v1/users", db, router).UsersRoute()
	routes.NewCustomersRoute("/api/v1/customers", db, router).CustomersRoute()
	routes.NewRolesRoute("/api/v1/roles", db, router).RolesRoute()
	routes.NewCategoriesRoute("/api/v1/categories", db, router).CategoriesRoute()

	httpErr := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("PORT")), router)
	if httpErr != nil {
		logrus.Errorf("http server not listening: %v", httpErr)
	}
}

func SetupDatabase() *sqlx.DB {
	db := configs.Connection("postgres")
	err := db.Ping()

	if err != nil {
		defer db.Close()
		logrus.Errorf("database not connected: %v", err)
	}

	if viper.GetString("GO_ENV") == "development" {
		logrus.Info("database is connected")
	}

	return db
}

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	if viper.GetString("GO_ENV") == "development" {
		logrus.Info("server is running on port: " + viper.GetString("PORT"))
	}

	return router
}
