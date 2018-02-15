package osin

import (
	"encoding/base64"

	uuid "github.com/satori/go.uuid"
)

// AuthorizeTokenGenDefault is the default authorization token generator
type AuthorizeTokenGenDefault struct {
}

// GenerateAuthorizeToken generates a base64-encoded UUID code
func (a *AuthorizeTokenGenDefault) GenerateAuthorizeToken(data *AuthorizeData) (ret string, err error) {
	token, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString([]byte(token.String())), nil
}

// AccessTokenGenDefault is the default authorization token generator
type AccessTokenGenDefault struct {
}

// GenerateAccessToken generates base64-encoded UUID access and refresh tokens
func (a *AccessTokenGenDefault) GenerateAccessToken(data *AccessData, generaterefresh bool) (accesstoken string, refreshtoken string, err error) {
	token, err := uuid.NewV4()
	if err != nil {
		return "", "", err
	}
	accesstoken = base64.RawURLEncoding.EncodeToString([]byte(token.String()))

	if generaterefresh {
		rtoken, err := uuid.NewV4()
		if err != nil {
			return "", "", err
		}
		refreshtoken = base64.RawURLEncoding.EncodeToString([]byte(rtoken.String()))
	}
	return
}
