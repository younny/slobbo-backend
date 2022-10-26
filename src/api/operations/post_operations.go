package operations

import (
	"net/http"

	m "github.com/younny/slobbo-backend/src/middleware"
	"github.com/younny/slobbo-backend/src/types"

	"github.com/go-chi/render"
)

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(m.PageIDKey)
	if err := render.Render(w, r, server.DB.GetPosts(pageID.(int))); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(m.PostCtxKey).(*types.Post)

	if err := render.Render(w, r, post); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	postRequest := &types.PostRequest{}

	if err := render.Bind(r, postRequest); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := postRequest.Validate(); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	newPost := types.Post{
		Title:     postRequest.Title,
		SubTitle:  postRequest.SubTitle,
		Body:      postRequest.Body,
		Author:    postRequest.Author,
		Category:  postRequest.Category,
		Thumbnail: postRequest.Thumbnail,
	}

	if err := server.DB.CreatePost(&newPost); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, &newPost); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(m.PostCtxKey).(*types.Post)

	if err := render.Bind(r, post); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
	}

	if err := server.DB.UpdatePost(post); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, post); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
	}
}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(m.PostCtxKey).(*types.Post)

	if err := server.DB.DeletePost(post.ID); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}
}
