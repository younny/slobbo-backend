package types

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	SubTitle  string    `json:"sub_title"`
	Body      string    `json:"body" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Category  uint      `json:"category"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}

func (p *Post) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *Post) Bind(r *http.Request) error {
	return nil
}

type PostList struct {
	Items      []*Post `json:"items"`
	NextPageID uint    `json:"next_page_id,omitempty"`
}

func (p *PostList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PostRequest struct {
	Title     string `json:"title" gorm:"not null"`
	SubTitle  string `json:"sub_title"`
	Body      string `json:"body" gorm:"not null"`
	Author    string `json:"author" gorm:"not null"`
	Category  uint   `json:"category"`
	Thumbnail string `json:"thumbnail"`
}

func (p *PostRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *PostRequest) Bind(r *http.Request) error {
	return nil
}

type ErrResponse struct {
	Err            error  `json:"-"`
	HttpStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	StatusCode     int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}

func ErrInvalidRequst(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HttpStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HttpStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func ErrNotFound() render.Renderer {
	return &ErrResponse{
		HttpStatusCode: http.StatusNotFound,
		StatusText:     "Resources not found.",
	}
}
