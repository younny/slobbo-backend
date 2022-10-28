package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/younny/slobbo-backend/src/db"
	"github.com/younny/slobbo-backend/src/types"
)

type (
	CustomKey string
)

const (
	UserCtxKey     CustomKey = "user"
	PostCtxKey     CustomKey = "post"
	WorkshopCtxKey CustomKey = "workshop"
)

var DBClient db.ClientInterface

func SetDBClient(c db.ClientInterface) {
	DBClient = c
}

func User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *types.User

		if id := chi.URLParam(r, "id"); id != "" {
			userId, err := strconv.Atoi(id)
			if err != nil {
				_ = render.Render(w, r, types.ErrInvalidRequst(err))
				return
			}

			user = DBClient.GetUserByID(uint(userId))
		} else {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		if user == nil {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

func Workshop(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var workshop *types.Workshop

		if id := chi.URLParam(r, "id"); id != "" {
			workshopID, err := strconv.Atoi(id)
			if err != nil {
				_ = render.Render(w, r, types.ErrInvalidRequst(err))
			}
			workshop = DBClient.GetWorkshopByID(uint(workshopID))
		} else {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		if workshop == nil {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), WorkshopCtxKey, workshop)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
