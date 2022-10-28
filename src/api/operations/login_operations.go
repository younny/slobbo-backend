package operations

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/younny/slobbo-backend/src/auth"
	"github.com/younny/slobbo-backend/src/types"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}

	if err := render.Bind(r, user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	user.Prepare()
	if err := user.Validate("login"); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	// todo token

}

func (server *Server) Signin(email, password string) error {
	user := server.DB.GetUserByEmail(email)

	if user == nil {
		return errors.New("User not found")
	}

	if err := types.VerifyPassword(user.Password, password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return auth.CreateToken(user.ID)
}
