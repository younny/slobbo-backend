package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/younny/slobbo-backend/src/api/operations"
	l "github.com/younny/slobbo-backend/src/log"
	"go.uber.org/zap"
)

var server = operations.Server{}
var (
	version string
	addr    string
)

func init() {
	pflag.StringVarP(&addr, "address", "a", ":8080", "the address for the api to listen on. host and port seperated by ':'")
	pflag.Parse()

	if err := godotenv.Load(); err != nil {
		l.Log.Error("Couldn't get env", zap.Error(err))
	}
}

func Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	dbUri := fmt.Sprintf("host=%s port=%s sslmode=disable user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	server.Initialize(dbUri)

	server.Run(addr)

	l.Log.Info("Ready to serve reqeusts on " + addr)
	<-c
	l.Log.Info("Gracefully shutting down")
	os.Exit(0)
}
