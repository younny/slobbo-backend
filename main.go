package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"slobbo/src/api"
	"slobbo/src/db"
)

var (
	version string
	addr    string
)

func init() {
	pflag.StringVarP(&addr, "address", "a", ":8080", "the address for the api to listen on. host and port seperated by ':'")
	pflag.Parse()
}

func main() {

	// exit on keyboard interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// logger
	log, _ := zap.NewProduction(zap.WithCaller(false))
	defer func() {
		_ = log.Sync()
	}()

	log.Info("starting up API...", zap.String("version", version))

	// Set database
	dbClient := &db.Client{}
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "postgresdb", "5432", "younny", "slobbo", "SlobboDibo")
	if err := dbClient.Connect(dbUri); err != nil {
		log.Error("Couldn't connect to database", zap.Error(err))
		os.Exit(1)
	}

	// Start api server
	r := api.GetRouter(log, dbClient)
	go func() {
		if err := http.ListenAndServe(addr, r); err != nil {
			log.Error("Failed to start server", zap.Error(err))
			return
		}
	}()

	log.Info("Ready to serve reqeusts on " + addr)
	<-c
	log.Info("Gracefully shutting down")
	os.Exit(0)
}
