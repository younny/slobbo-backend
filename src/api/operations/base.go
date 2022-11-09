package operations

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/younny/slobbo-backend/src/db"
	l "github.com/younny/slobbo-backend/src/log"
	m "github.com/younny/slobbo-backend/src/middleware"
)

type Server struct {
	DB     db.ClientInterface
	Router *chi.Mux
}

func (server *Server) Initialize(DbUri string) {
	dbClient := &db.Client{}

	if err := dbClient.Connect(DbUri); err != nil {
		l.Log.Error("Couldn't connect to database", zap.Error(err))
		os.Exit(1)
	}
	server.Set(dbClient)
}

func (server *Server) Run(addr string) {
	go func() {
		if err := http.ListenAndServe(addr, server.Router); err != nil {
			l.Log.Error("Failed to start server", zap.Error(err))
			return
		}
	}()
}

func (server *Server) Set(dbClient db.ClientInterface) {
	server.Router = getRouter()
	server.DB = dbClient
	m.SetDBClient(dbClient)
	server.initializeRoutes()
}

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID,
		middleware.Recoverer,
		middleware.RealIP,
		middleware.Logger,
		middleware.Timeout(60*time.Second),
	)

	return r
}
