package models

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Session struct {
	Token string `json:"token"`
	Expiration time.Time `json:"expiration"`
	UserID string `json:"user_id"`
}

func NewSession(userID string) *Session {
	token := uuid.New().String()
	expiration := time.Now()

	return &Session{
		Token: token,
		Expiration: expiration,
		UserID: userID,
	}
}

func (s * Session) GetCookie() *http.Cookie {
	return &http.Cookie{
		Name:       "Test_auth_service",
		Value:      s.Token,
		MaxAge:    3600 * 12,
		HttpOnly:   true,
		SameSite: http.SameSiteNoneMode,
	}
}