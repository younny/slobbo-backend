package operations

import (
	"github.com/go-chi/chi/v5"
	m "github.com/younny/slobbo-backend/src/middleware"
)

func (s *Server) initializeRoutes() {
	s.Router.Route("/login", func(r chi.Router) {
		r.Post("/", s.Login)
	})
	s.Router.With(m.Auth).Route("/users", func(r chi.Router) {
		r.Get("/", s.GetUsers)
		r.Post("/", s.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.User)
			r.Get("/", GetUserByID)
			r.Patch("/", s.UpdateUser)
			r.Delete("/", s.DeleteUser)
		})
	})
	s.Router.Route("/about", func(r chi.Router) {
		r.Get("/", s.GetAbout)
		r.With(m.Auth).Post("/", s.CreateAbout)
		r.With(m.Auth).Patch("/", s.UpdateAbout)
	})
	s.Router.Route("/workshops", func(r chi.Router) {
		r.With(m.Pagination).Get("/", s.GetWorkshops)
		r.With(m.Auth).Post("/", s.CreateWorkshop)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Workshop)
			r.Get("/", GetWorkshop)
			r.With(m.Auth).Patch("/", s.UpdateWorkshop)
			r.With(m.Auth).Delete("/", s.DeleteWorkshop)
		})
	})
	s.Router.Route("/posts", func(r chi.Router) {
		r.With(m.Pagination).Get("/", s.GetPosts)
		r.With(m.Auth).Post("/", s.CreatePost)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Post)
			r.Get("/", GetPost)
			r.With(m.Auth).Patch("/", s.UpdatePost)
			r.With(m.Auth).Delete("/", s.DeletePost)
		})
	})
}
