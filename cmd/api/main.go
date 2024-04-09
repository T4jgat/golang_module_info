package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	// config struct instance
	var cfg config

	// defining command line flags (if flags was not inserted, server uses default values)
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// logger instance
	// writes logs with prefixes - date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// application instance
	app := &application{
		config: cfg,
		logger: logger,
	}

	// server instance
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// server staring log message
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	// error logs
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
