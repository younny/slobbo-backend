package operations

import (
	"github.com/go-chi/chi/v5"
	m "github.com/younny/slobbo-backend/src/middleware"
)

func (s *Server) initializeRoutes() {
	s.Router.Route("/workshops", func(r chi.Router) {
		r.With(m.Pagination).Get("/", s.GetWorkshops)
		r.Post("/", s.CreateWorkshop)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Workshop)
			r.Get("/", GetWorkshop)
			r.Patch("/", s.UpdateWorkshop)
			r.Delete("/", s.DeleteWorkshop)
		})
	})
	s.Router.Route("/posts", func(r chi.Router) {
		r.With(m.Pagination).Get("/", s.GetPosts)
		r.Post("/", s.CreatePost)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Post)
			r.Get("/", GetPost)
			r.Patch("/", s.UpdatePost)
			r.Delete("/", s.DeletePost)
		})
	})
}
