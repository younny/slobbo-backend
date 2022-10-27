package types

import "net/http"

type About struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `gorm:"not null" json:"title"`
	SubTitle string    `json:"sub_title"`
	Body1    string    `json:"body_1"`
	Body2    string    `json:"body_2"`
	Contacts *Contacts `gorm:"embedded;embedded_prefix:contacts_" json:"contacts"`
}

func (about *About) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (about *About) Bind(r *http.Request) error {
	return nil
}

type Contacts struct {
	Email  string `json:"email"`
	Github string `json:"github"`
}
