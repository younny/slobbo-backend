package api

import (
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"slobbo/src/db"
	m "slobbo/src/middleware"
)

var DBClient db.ClientInterface

func SetDBClient(c db.ClientInterface) {
	DBClient = c
	m.SetDBClient(DBClient)
}

func GetRouter(log *zap.Logger, dbClient db.ClientInterface) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	SetDBClient(dbClient)

	if log != nil {
		r.Use(m.SetLogger(log))
	}
	buildTree(r)

	return r
}

func buildTree(r *chi.Mux) {
	r.Route("/posts", func(r chi.Router) {
		r.With(m.Pagination).Get("/", GetPosts)
		r.Post("/", CreatePost)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Post)
			r.Get("/", GetPost)
			r.Patch("/", UpdatePost)
			r.Delete("/", DeletePost)
		})
	})
}
