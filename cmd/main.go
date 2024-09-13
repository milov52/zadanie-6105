package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/common/server"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/repository/pgrepo"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/services"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/transport/httpserver"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/pkg/pg"
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

	// create repositories
	userRepo := pgrepo.NewUserRepo(pgDB)
	tenderRepo := pgrepo.NewTenderRepo(pgDB)
	bidRepo := pgrepo.NewBidRepo(pgDB)

	userService := services.NewUserService(userRepo)
	tenderService := services.NewTenderService(tenderRepo)
	bidService := services.NewBidService(bidRepo)

	// create http server with application injected
	httpServer := httpserver.NewHttpServer(userService, tenderService, bidService)

	// create http router
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		server.RespondOK("ok", w, r)
	}).Methods(http.MethodGet)

	tenderRouter := apiRouter.PathPrefix("/tenders").Subrouter()
	tenderRouter.HandleFunc("", httpServer.GetTenders).Methods(http.MethodGet)
	tenderRouter.HandleFunc("/new", httpServer.CreateTender).Methods(http.MethodPost)
	tenderRouter.HandleFunc("/my", httpServer.GetUserTenders).Methods(http.MethodGet)
	tenderRouter.HandleFunc("/{tenderId}/status", httpServer.GetTenderStatus).Methods(http.MethodGet)
	tenderRouter.HandleFunc("/{tenderId}/status", httpServer.UpdateTenderStatus).Methods(http.MethodPut)
	tenderRouter.HandleFunc("/{tenderId}/edit", httpServer.UpdateTender).Methods(http.MethodPatch)
	tenderRouter.HandleFunc("/{tenderId}/rollback/{version}", httpServer.RollbackVersion).Methods(http.MethodPut)

	bidsRouter := apiRouter.PathPrefix("/bids").Subrouter()
	bidsRouter.HandleFunc("/new", httpServer.CreateBid).Methods(http.MethodPost)
	bidsRouter.HandleFunc("/my", httpServer.GetUserBids).Methods(http.MethodGet)
	bidsRouter.HandleFunc("/{tenderID}/list", httpServer.GetTenderBids).Methods(http.MethodGet)
	bidsRouter.HandleFunc("/{bidID}/status", httpServer.GetBidStatus).Methods(http.MethodGet)
	bidsRouter.HandleFunc("/{bidID}/status", httpServer.UpdateBidStatus).Methods(http.MethodPut)
	bidsRouter.HandleFunc("/{bidId}/edit", httpServer.UpdateBid).Methods(http.MethodPatch)
	bidsRouter.HandleFunc("/{bidId}/submit_decision", httpServer.SubmitDecision).Methods(http.MethodPut)
	bidsRouter.HandleFunc("/{bidId}/feedback", httpServer.BidFeedback).Methods(http.MethodPut)
	bidsRouter.HandleFunc("/{bidId}/rollback/{version}", httpServer.RollbackBidVersion).Methods(http.MethodPut)
	bidsRouter.HandleFunc("/{tenderId}/reviews", httpServer.GetReviews).Methods(http.MethodGet)

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
