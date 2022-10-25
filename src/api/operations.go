package api

import (
	"net/http"
	m "slobbo/src/middleware"
	"slobbo/src/types"

	"github.com/go-chi/render"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(m.PageIDKey)
	if err := render.Render(w, r, DBClient.GetPosts(pageID.(int))); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(m.PostCtxKey).(*types.Post)
	println("post (id=%d, title=%s)", post.ID, post.Title)
	if err := render.Render(w, r, post); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	request := &types.PostRequest{}

	if err := render.Bind(r, request); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	newPost := types.Post{
		Title:     request.Title,
		SubTitle:  request.SubTitle,
		Body:      request.Body,
		Author:    request.Author,
		Category:  request.Category,
		Thumbnail: request.Thumbnail,
	}

	if err := DBClient.CreatePost(&newPost); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, &newPost); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(m.PostCtxKey).(*types.Post)

	if err := render.Bind(r, post); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
	}

	if err := DBClient.UpdatePost(post); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, post); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(m.PostCtxKey).(*types.Post)

	if err := DBClient.DeletePost(post.ID); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}
}
