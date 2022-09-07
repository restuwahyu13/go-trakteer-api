package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	routes.UsersRoute("/api/v1/users", db, router)

	httpErr := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("PORT")), router)
	if httpErr != nil {
		log.Fatalf("http server not listening: %v", httpErr)
	}
}

func SetupDatabase() *sqlx.DB {
	db := configs.Connection("postgres")
	err := db.Ping()

	if err != nil {
		defer db.Close()
		log.Fatalf("database not connected: %v", err)
	}

	if viper.GetString("GO_ENV") == "development" {
		log.Print("database is connected")
	}

	return db
}

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	if viper.GetString("GO_ENV") == "development" {
		log.Println("server is running on port: " + viper.GetString("PORT"))
	}

	return router
}
