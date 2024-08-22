package presentation

import (
	"errors"
	"strings"

	"github.com/wizard-corp/api-gateway/domain"
)

func NewLoginController(
	uc domain.LoginUsecase,
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int,
) (*domain.LoginResponse, error) {
	errs := domain.IsLoginValid(email, password)
	if len(errs) > 0 {
		return nil, errors.New(domain.INVALID_SCHEMA + "\n" + strings.Join(errs, "\n"))
	}

	loginResponse, err := uc.NewLogin(
		email,
		password,
		accessTokenSecret,
		accessTokenExpiryHour,
		refreshTokenSecret,
		refreshTokenExpiryHour,
	)
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}
