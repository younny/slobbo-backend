package middleware

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/younny/slobbo-backend/src/auth"
	"github.com/younny/slobbo-backend/src/types"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			_ = render.Render(w, r, types.ErrNotAuthorised(err))
			return
		}
		next.ServeHTTP(w, r)
	})
}
