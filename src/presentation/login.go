package presentation

import (
	"errors"
	"strings"
	"time"

	"github.com/wizard-corp/api-gateway/src/application"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/mymongo"
)

type LoginController struct {
	uc *application.LoginUsecase
}

func NewLoginController(timeout time.Duration, app *bootstrap.App) *LoginController {
	repo := mymongo.NewUserRepository(app.Mongo, timeout)
	uc := application.LoginUsecase{LoginRepo: repo}
	return &LoginController{&uc}
}

func (lc *LoginController) NewLogin(
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*domain.JwtTokenResponse, error) {
	login := domain.Login{
		Email:    email,
		Password: password,
		JwtToken: domain.JwtToken{
			AccessTokenSecret:      accessTokenSecret,
			AccessTokenExpiryHour:  accessTokenExpiryHour,
			RefreshTokenSecret:     refreshTokenSecret,
			RefreshTokenExpiryHour: refreshTokenExpiryHour}}
	errs := login.IsLoginValid()
	if len(errs) > 0 {
		return nil, errors.New(domain.INVALID_SCHEMA + "\n" + strings.Join(errs, "\n"))
	}

	loginResponse, err := lc.uc.NewLogin(&login)
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}
