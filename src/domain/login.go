package domain

import (
	"strings"
)

type Login struct {
	Email    string
	Password string
	JwtToken
}

type LoginRepository interface {
	GetUserByEmail(email string) (User, error)
}

func (l *Login) IsLoginValid() []string {
	var errors []string

	if l.Email == "" {
		errors = append(errors, IS_EMPTY)
	}

	if l.Password == "" {
		errors = append(errors, IS_EMPTY)
	}

	if IsValidEmail(l.Email) {
		errors = append(errors, INVALID_FORMAT)
	}

	return errors
}

func NewLogin(
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*Login, error) {
	login := Login{
		Email:    email,
		Password: password,
		JwtToken: JwtToken{
			AccessTokenSecret:      accessTokenSecret,
			AccessTokenExpiryHour:  accessTokenExpiryHour,
			RefreshTokenSecret:     refreshTokenSecret,
			RefreshTokenExpiryHour: refreshTokenExpiryHour}}
	errs := login.IsLoginValid()
	if len(errs) > 0 {
		return nil, NewDomainError(INVALID_SCHEMA, strings.Join(errs, "\n"))
	}
	return &login, nil
}
