package types

import (
	"errors"
	"net/http"
	"time"
)

type Workshop struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Name      string        `gorm:"not null" json:"name"`
	Details   string        `gorm:"not null" json:"details"`
	Organiser string        `gorm:"not null" json:"organiser"`
	Location  string        `gorm:"not null" json:"location"`
	Datetime  time.Time     `gorm:"not null" json:"datetime"`
	Duration  time.Duration `gorm:"not null" json:"duration"`
	Capacity  int           `gorm:"not null" json:"capacity"`
	Price     string        `gorm:"not null" json:"price"`
	CreatedAt time.Time     `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

func (p *Workshop) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *Workshop) Bind(r *http.Request) error {
	return nil
}

func (p *Workshop) Validate() error {
	if p.Name == "" {
		return errors.New("Name required")
	}
	if p.Details == "" {
		return errors.New("Details required")
	}
	if p.Organiser == "" {
		return errors.New("Organiser required")
	}
	if p.Location == "" {
		return errors.New("Location required")
	}
	if p.Datetime.String() == "" {
		return errors.New("Datetime required")
	}
	if p.Duration == 0 {
		return errors.New("Duration is more than zero")
	}
	if p.Capacity == 0 {
		return errors.New("Capacity is more than zero")
	}
	if p.Price == "" {
		return errors.New("Price required")
	}

	return nil
}

type WorkshopList struct {
	Items      []*Workshop `json:"items"`
	NextPageID uint        `json:"next_page_id,omitempty"`
}

func (wl *WorkshopList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type WorkshopRequest struct {
	Name      string        `gorm:"not null" json:"name"`
	Details   string        `gorm:"not null" json:"details"`
	Organiser string        `gorm:"not null" json:"organiser"`
	Location  string        `gorm:"not null" json:"location"`
	Datetime  time.Time     `gorm:"not null" json:"datetime"`
	Duration  time.Duration `gorm:"not null" json:"duration"`
	Capacity  int           `gorm:"not null" json:"capacity"`
	Price     string        `gorm:"not null" json:"price"`
}

func (p *WorkshopRequest) Validate() error {
	if p.Name == "" {
		return errors.New("Name required")
	}
	if p.Details == "" {
		return errors.New("Details required")
	}
	if p.Organiser == "" {
		return errors.New("Organiser required")
	}
	if p.Location == "" {
		return errors.New("Location required")
	}
	if p.Datetime.String() == "" {
		return errors.New("Datetime required")
	}
	if p.Duration == 0 {
		return errors.New("Duration is more than zero")
	}
	if p.Capacity == 0 {
		return errors.New("Capacity is more than zero")
	}
	if p.Price == "" {
		return errors.New("Price required")
	}

	return nil
}

func (p *WorkshopRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *WorkshopRequest) Bind(r *http.Request) error {
	return nil
}
