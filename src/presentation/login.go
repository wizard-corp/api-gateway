package presentation

import (
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

func (lc *LoginController) Login(
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*domain.JwtTokenResponse, error) {
	login, err := domain.NewLogin(
		email,
		password,
		accessTokenSecret,
		accessTokenExpiryHour,
		refreshTokenSecret,
		refreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	loginResponse, err := lc.uc.Login(login)
	if err != nil {
		return nil, err
	}

	return loginResponse, nil
}
