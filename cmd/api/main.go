package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/T4jgat/module_info/internal/data"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

// Dependency injection
type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {

	// config struct instance
	var cfg config

	// defining command line flags (if flags was not inserted, server uses default values)
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DSN"), "PostgreSQL DSN")

	flag.Parse()

	// logger instance
	// writes logs with prefixes - date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Printf("database connection pool established")

	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal(err, nil)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://C:/Home/code/golangprojects/adv-prog-assign-1/migrations",
		"postgres",
		migrationDriver,
	)
	if err != nil {
		logger.Fatal(err, nil)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal(err, nil)
	}

	logger.Printf("database migrations applied")

	// application instance
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	// server instance
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	} // server staring log message
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	// error logs
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
