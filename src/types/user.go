package types

import (
	"errors"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}

type UserList struct {
	Items []*User `json:"items"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Username required")
		}
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if u.Username == "" {
			return errors.New("Username required")
		}
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *User) Bind(r *http.Request) error {
	return nil
}

func (ul *UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
