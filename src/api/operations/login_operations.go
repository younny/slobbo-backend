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

	token, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		if err.Error() == "UserNotFound" {
			_ = render.Render(w, r, types.ErrNotFound())
		}
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	render.Render(w, r, &types.TokenResponse{Token: token})
}

func (server *Server) SignIn(email, password string) (string, error) {
	user := server.DB.GetUserByEmail(email)

	if user == nil {
		return "", errors.New("UserNotFound")
	}

	if err := types.VerifyPassword(user.Password, password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
