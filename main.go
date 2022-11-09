package main

import (
	"os"

	"github.com/younny/slobbo-backend/src/api"
	l "github.com/younny/slobbo-backend/src/log"
)

var (
	version string
	addr    string
)

func main() {
	env := os.Args[1:][0]
	l.SetLogger()

	api.Run(env)
}
