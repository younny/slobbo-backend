package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"github.com/younny/slobbo-backend/src/types"
)

const (
	PageIDKey CustomKey = "page_id"
)

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PageID := r.URL.Query().Get(string(PageIDKey))
		intPageID := 0

		var err error
		if PageID != "" {
			intPageID, err = strconv.Atoi(PageID)
			if err != nil {
				_ = render.Render(w, r, types.ErrInvalidRequst(fmt.Errorf("couldn't read %s: %w", PageIDKey, err)))
				return
			}
		}
		ctx := context.WithValue(r.Context(), PageIDKey, intPageID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
