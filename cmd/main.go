package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	stdlog "log"

	log "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/suprageeks/tech-task/middleware"
	"github.com/suprageeks/tech-task/pkg/handlers"
)

func main() {
	// Capture Ctrl-C
	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	//--------------logging middleware---------------------
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	//---------------------------------------------------
	router := mux.NewRouter()
	handler := handlers.NewBankAndAccountsHandlerInstance()
	fmt.Println("started listening on port : ", 9294)
	router.Handle("/bank", http.HandlerFunc(handler.CreateBank)).Methods("POST")
	router.Handle("/bank", http.HandlerFunc(handler.GetAllBanks)).Methods("GET")
	router.Handle("/bank/{id}", http.HandlerFunc(handler.GetBankDetails)).Methods("GET")
	router.Handle("/bank/{id}", http.HandlerFunc(handler.UpdateBankDetails)).Methods("PATCH")
	router.Handle("/bank/{id}", http.HandlerFunc(handler.RemoveBank)).Methods("DELETE")

	router.Handle("/account", http.HandlerFunc(handler.CreateAccount)).Methods("POST")
	router.Handle("/account/{id}", http.HandlerFunc(handler.GetAccountDetails)).Methods("GET")

	//-------------------------prometheous endpoints--------------------
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static/")))
	router.Path("/prometheus").Handler(promhttp.Handler())
	//--------------------------------------------------------------------

	loggedRouter := loggingMiddleware(router)
	if err := http.ListenAndServe(":9294", loggedRouter); err != nil {
		logger.Log("status", "fatal", "err", err)
	}

	//-------------capturing the ctrl + c event----------------------
	select {
	case <-c:
		stdlog.Println("cancel operation")
		cancel()

	case <-ctx.Done():
		time.Sleep(600 * time.Millisecond)
	}

	stdlog.Println("done")
}
