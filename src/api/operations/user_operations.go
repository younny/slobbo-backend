package operations

import (
	"net/http"

	"github.com/go-chi/render"
	m "github.com/younny/slobbo-backend/src/middleware"
	"github.com/younny/slobbo-backend/src/types"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	if err := render.Render(w, r, server.DB.GetUsers()); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(m.UserCtxKey).(*types.User)

	if err := render.Render(w, r, user); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &types.User{}

	if err := render.Bind(r, newUser); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	newUser.Prepare()

	if err := newUser.Validate(""); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := server.DB.CreateUser(newUser); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, newUser); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}

	if err := render.Bind(r, user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	user.Prepare()

	if err := user.Validate(""); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := server.DB.UpdateUser(user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, user); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(m.UserCtxKey).(*types.User)

	if err := server.DB.DeleteUser(user.ID); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}
}
