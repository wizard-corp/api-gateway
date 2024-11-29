package domain

import (
	"strings"
)

type RefreshToken struct {
	RefreshToken string
	JwtToken
}

type RefreshTokenRepository interface {
	GetUserByID(id string) (User, error)
}

func (rtc *RefreshToken) IsRefreshTokenValid() []string {
	var errors []string

	if rtc.RefreshToken == "" {
		errors = append(errors, IS_EMPTY)
	}

	return errors
}

func NewRefreshToken(
	refreshToken string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*RefreshToken, error) {
	rtc := RefreshToken{
		RefreshToken: refreshToken,
		JwtToken: JwtToken{
			AccessTokenSecret:      accessTokenSecret,
			AccessTokenExpiryHour:  accessTokenExpiryHour,
			RefreshTokenSecret:     refreshTokenSecret,
			RefreshTokenExpiryHour: refreshTokenExpiryHour}}
	errs := rtc.IsRefreshTokenValid()
	if len(errs) > 0 {
		return nil, NewDomainError(INVALID_SCHEMA, strings.Join(errs, "\n"))
	}
	return &rtc, nil
}
