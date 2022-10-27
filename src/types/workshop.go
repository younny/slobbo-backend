package types

import (
	"errors"
	"net/http"
	"time"
)

type Organiser struct {
	Name string `gorm:"not null" json:"name"`
	Link string `json:"link"`
}

type Duration struct {
	StartDate    string `gorm:"not null" json:"start_date"`
	EndDate      string `gorm:"not null" json:"end_date"`
	TotalInHours int    `gorm:"not null" json:"total_hrs"`
}

type Price struct {
	Amount   int    `gorm:"not null" json:"amount"`
	Currency string `gorm:"default:'Euro'" json:"currency"`
}

type Workshop struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Organiser   *Organiser `gorm:"not null;embedded;embedded_prefix:org_" json:"organiser"`
	Location    string     `gorm:"not null" json:"location"`
	Duration    *Duration  `gorm:"not null;embedded;embedded_prefix:dur_" json:"duration"`
	Capacity    int        `gorm:"not null" json:"capacity"`
	Price       *Price     `gorm:"not null;embedded;embedded_prefix:price_" json:"price"`
	Thumbnail   string     `json:"thumbnail"`
	CreatedAt   time.Time  `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

func (ws *Workshop) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ws *Workshop) Bind(r *http.Request) error {
	return nil
}

func (ws *Workshop) Validate() error {
	if ws.Name == "" {
		return errors.New("Name required")
	}

	if ws.Description == "" {
		return errors.New("Description required")
	}

	if ws.Organiser.Name == "" {
		return errors.New("Organiser name required")
	}

	if ws.Location == "" {
		return errors.New("Location required")
	}

	if _, err := time.Parse(time.RFC3339, ws.Duration.StartDate); err != nil {
		return errors.New("Duration 'start_date' format should be " + time.RFC3339)
	}

	if _, err := time.Parse(time.RFC3339, ws.Duration.EndDate); err != nil {
		return errors.New("Duration 'end_date' format should be " + time.RFC3339)
	}

	if ws.Duration.TotalInHours <= 0 {
		return errors.New("Duration 'total_hrs' more than zero")
	}

	if ws.Capacity <= 0 {
		return errors.New("Capacity is more than zero")
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
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Organiser   *Organiser `gorm:"not null" json:"organiser"`
	Location    string     `gorm:"not null" json:"location"`
	Duration    *Duration  `gorm:"not null" json:"duration"`
	Capacity    int        `gorm:"not null" json:"capacity"`
	Price       *Price     `gorm:"not null" json:"price"`
	Thumbnail   string     `json:"thumbnail"`
}

func (wr *WorkshopRequest) Validate() error {
	if wr.Name == "" {
		return errors.New("Name required")
	}

	if wr.Description == "" {
		return errors.New("Details required")
	}

	if wr.Organiser == nil {
		return errors.New("Organiser required")
	}

	if wr.Location == "" {
		return errors.New("Location required")
	}

	_, err := time.Parse(time.RFC3339, wr.Duration.StartDate)

	if err != nil {
		return errors.New("Duration 'start_date' format should be " + time.RFC3339 + "(error: " + err.Error() + ")")
	}

	if _, err := time.Parse(time.RFC3339, wr.Duration.EndDate); err != nil {
		return errors.New("Duration 'end_date' format should be " + time.RFC3339 + "(error: " + err.Error() + ")")
	}

	if wr.Duration.TotalInHours <= 0 {
		return errors.New("Duration 'total_hrs' more than zero")
	}

	if wr.Capacity == 0 {
		return errors.New("Capacity is more than zero")
	}

	return nil
}

func (wr *WorkshopRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wr *WorkshopRequest) Bind(r *http.Request) error {
	return nil
}
