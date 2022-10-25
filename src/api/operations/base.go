package operations

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/younny/slobbo-backend/src/db"
	l "github.com/younny/slobbo-backend/src/log"
	m "github.com/younny/slobbo-backend/src/middleware"
)

type Server struct {
	DB     *db.Client
	Router *chi.Mux
}

func (server *Server) Initialize(DbUri string) {
	dbClient := &db.Client{}

	if err := dbClient.Connect(DbUri); err != nil {
		l.Log.Error("Couldn't connect to database", zap.Error(err))
		os.Exit(1)
	}
}

func (server *Server) Run(addr string) {
	dbClient := &db.Client{}

	r := GetRouter(dbClient)
	server.DB = dbClient
	server.Router = r
	m.SetDBClient(dbClient)

	server.initializeRoutes()

	go func() {
		if err := http.ListenAndServe(addr, r); err != nil {
			l.Log.Error("Failed to start server", zap.Error(err))
			return
		}
	}()
}

func GetRouter(dbClient db.ClientInterface) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)

	if l.Log != nil {
		r.Use(m.SetLogger(l.Log))
	}

	return r
}
