package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
)

type User struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

func NewUser(login string, email string, password string, phone string) (*User, error) {
	user := &User{
		Login:    login,
		Email:    email,
		Password: password,
		Phone:    phone,
	}

	if user.Login == "" {
		return nil, LoginFormatError
	}

	if !user.ValidateEmail() {
		return nil, EmailFormatError
	}

	if !user.ValidatePhone() {
		return nil, PhoneFormatError
	}

	if !user.ValidatePassword() {
		return nil, PasswordFormatError
	}

	return user, nil
}

func (u *User) ValidatePassword() bool {
	if u.Password == "" {
		return false
	}

	pattern := `^[a-zA-Z]\w{3,14}$`

	match, err := regexp.Match(pattern, []byte(u.Password))
	if err != nil {
		log.Print(err)
		return false
	}

	if !match {
		return false
	}

	return true
}

func (u *User) ValidatePhone() bool {
	if u.Phone == "" {

	}

	pattern := `^\d+.{10,13}$`

	match, err := regexp.Match(pattern, []byte(u.Phone))
	if err != nil {
		log.Print(err)
		return false
	}

	if !match {
		return false
	}

	return true
}

func (u *User) ValidateEmail() bool {
	if u.Email == "" {
		return false
	}

	pattern := `^\w+@\w+\.\w+$`

	match, err := regexp.Match(pattern, []byte(u.Email))
	if err != nil {
		log.Print(err)
		return false
	}

	if !match {
		return false
	}

	return true
}

func (u *User) HashPassword() {
	h := sha256.New()
	h.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(h.Sum(nil))
}

func (u *User) VerifyPassword(pass string) bool {
	h := sha256.New()
	h.Write([]byte(pass))
	fmt.Printf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil)) == u.Password
}