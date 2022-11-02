package operations

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/younny/slobbo-backend/src/types"
)

func (server *Server) GetAbout(w http.ResponseWriter, r *http.Request) {
	if err := render.Render(w, r, server.DB.GetAbout()); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) CreateAbout(w http.ResponseWriter, r *http.Request) {
	about := &types.About{}

	if err := render.Bind(r, about); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := server.DB.CreateAbout(about); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, about); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) UpdateAbout(w http.ResponseWriter, r *http.Request) {
	about := server.DB.GetAbout()
	if about == nil {
		_ = render.Render(w, r, types.ErrNotFound())
		return
	}

	if err := render.Bind(r, about); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := server.DB.UpdateAbout(about); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, about); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}
}
