package types

import (
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
	CreatedAt time.Time     `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time     `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}

func (p *Workshop) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *Workshop) Bind(r *http.Request) error {
	return nil
}
