package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/gorilla/mux"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/repository/pgrepo"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/transport/httpserver"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/pkg/pg"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatal()
	}
	os.Exit(0)
}

func run() error {
	// read config from env
	cfg := config.Read()

	pgDB, err := pg.Dial(cfg.DSN)
	if err != nil {
		return fmt.Errorf("pg.Dial failed: %w", err)
	}

	// run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations")
		if err := runMigrations(cfg.DSN, cfg.MigrationsPath); err != nil {
			return fmt.Errorf("running migrations failed: %w", err)
		}
	}

	// create repositories
	userRepo := pgrepo.NewUserRepo(pgDB)

	userService := services.NewUserService(userRepo)

	// create http server with application injected
	httpServer := httpserver.NewHttpServer(userService)

	// create http router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Tenders API v0.1"))
	}).Methods("GET")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)
	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")
	return nil
}

// runPgMigrations runs Postgres migrations
func runMigrations(dsn, path string) error {
	if path == "" {
		return errors.New("no migrations path provided")
	}
	if dsn == "" {
		return errors.New("no DSN provided")
	}

	m, err := migrate.New(path, dsn)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
