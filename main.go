package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/restuwahyu13/go-trakteer-api/configs"
	"github.com/restuwahyu13/go-trakteer-api/packages"
	"github.com/restuwahyu13/go-trakteer-api/routes"
)

func main() {
	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	}

	if _, ok := os.LookupEnv("GO_ENV"); !ok {
		err := packages.Viper()
		if err != nil {
			log.Fatalf(".env file not load: %v", err)
		}
	}

	db := SetupDatabase()
	router := SetupRouter()
	SetupMiddleware(router)

	routes.NewUsersRoute("/api/v1/users", db, router).UsersRoute()
	routes.NewCustomersRoute("/api/v1/customers", db, router).CustomersRoute()
	routes.NewRolesRoute("/api/v1/roles", db, router).RolesRoute()
	routes.NewCategoriesRoute("/api/v1/categories", db, router).CategoriesRoute()

	SetupGraceFullShutDown(router)
}

func SetupDatabase() *sqlx.DB {
	db := configs.Connection("postgres")
	err := db.Ping()

	if err != nil {
		defer db.Close()
		logrus.Errorf("database not connected: %v", err)
	}

	if viper.GetString("GO_ENV") == "development" {
		logrus.Info("database connected")
	}

	return db
}

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	if viper.GetString("GO_ENV") == "development" {
		logrus.Info("Server running on port: " + viper.GetString("PORT"))
	}

	return router
}

func SetupMiddleware(router *chi.Mux) {
	if viper.GetString("GO_ENV") != "production" {
		router.Use(middleware.Logger)
	}

	router.Use(middleware.Compress(gzip.BestCompression))
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.AllowContentEncoding("application/json"))
	router.Use(middleware.ThrottleWithOpts(middleware.ThrottleOpts{Limit: 5, BacklogLimit: 50, BacklogTimeout: time.Duration(5 * time.Minute)}))
	router.Use(middleware.NoCache)
	router.Use(middleware.CleanPath)
	router.Use(middleware.RequestID)
}

func SetupGraceFullShutDown(router *chi.Mux) {
	httpServer := http.Server{
		Addr:           fmt.Sprintf(":%s", viper.GetString("PORT")),
		ReadTimeout:    time.Duration(time.Second) * 60,
		WriteTimeout:   time.Duration(time.Second) * 30,
		IdleTimeout:    time.Duration(time.Second) * 120,
		MaxHeaderBytes: 3145728,
		Handler:        router,
	}

	go func() {
		httpErr := httpServer.ListenAndServe()
		if httpErr != nil {
			logrus.Errorf("Server not running: %v", httpErr)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	logrus.Info(fmt.Sprintf("Signal received: %s", <-c))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second)*10)
	defer cancel()

	httpServer.Shutdown(ctx)
	logrus.Info("Http server shutdown")
}
