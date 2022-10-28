package types

import "net/http"

type TokenResponse struct {
	Token string `json:"token"`
}

func (t *TokenResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
