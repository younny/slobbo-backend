package main

import (
	"github.com/younny/slobbo-backend/src/api"
	l "github.com/younny/slobbo-backend/src/log"
)

var (
	version string
	addr    string
)

func main() {
	l.SetLogger()

	api.Run()
}
