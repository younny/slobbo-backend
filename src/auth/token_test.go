package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthTokenSuccess(t *testing.T) {
	token, err := CreateToken(0)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	bodyStr := `"{}"`
	body := bytes.NewReader([]byte(bodyStr))
	r, _ := http.NewRequest(http.MethodPost, "", body)
	r.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	err = TokenValid(r)
	assert.NoError(t, err)
}

func TestAuthTokenFailure(t *testing.T) {
	token, err := CreateToken(0)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	bodyStr := `"{}"`
	body := bytes.NewReader([]byte(bodyStr))
	r, _ := http.NewRequest(http.MethodPost, "", body)
	r.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer ")},
	}

	err = TokenValid(r)
	assert.Error(t, err)

	token = "this is wrong token"
	r, _ = http.NewRequest(http.MethodPost, "", body)
	r.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	err = TokenValid(r)
	assert.Error(t, err)
}

func TestExtractTokenID(t *testing.T) {
	token, err := CreateToken(0)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	bodyStr := `"{}"`
	body := bytes.NewReader([]byte(bodyStr))
	r, _ := http.NewRequest(http.MethodPost, "", body)
	r.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	tokenId, err := ExtractTokenID(r)
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), tokenId)

}
