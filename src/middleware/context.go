package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"slobbo/src/db"
	"slobbo/src/types"
)

type (
	CustomKey string
)

const (
	PostCtxKey CustomKey = "post"
)

var DBClient db.ClientInterface

func SetDBClient(c db.ClientInterface) {
	DBClient = c
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post *types.Post

		if id := chi.URLParam(r, "id"); id != "" {
			postID, err := strconv.Atoi(id)
			if err != nil {
				_ = render.Render(w, r, types.ErrInvalidRequst(err))
				return
			}
			post = DBClient.GetPostByID(uint(postID))
		} else {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		if post == nil {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), PostCtxKey, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
