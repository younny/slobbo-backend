package types

import (
	"errors"
	"net/http"
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"size:255;not null"`
	SubTitle  string    `json:"sub_title" gorm:"size:255"`
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
	Title     string `gorm:"size:255;not null" json:"title"`
	SubTitle  string `gorm:"size:255" json:"sub_title"`
	Body      string `gorm:"not null" json:"body"`
	Author    string `gorm:"not null" json:"author"`
	Category  uint   `json:"category"`
	Thumbnail string `json:"thumbnail"`
}

func (p *PostRequest) Validate() error {

	if p.Title == "" {
		return errors.New("Title required")
	}
	if p.Body == "" {
		return errors.New("Body required")
	}
	if p.Author == "" {
		return errors.New("Author required")
	}
	return nil
}

func (p *PostRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *PostRequest) Bind(r *http.Request) error {
	return nil
}
