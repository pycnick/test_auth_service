package models

import "time"

type Session struct {
	Token string `json:"token"`
	Expiration time.Time `json:"expiration"`
}