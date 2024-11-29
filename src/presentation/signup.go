package presentation

import (
	"time"

	"github.com/wizard-corp/api-gateway/src/application"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/mymongo"
)

type SignupController struct {
	uc *application.SignupUsecase
}

func NewSignupController(timeout time.Duration, app *bootstrap.App) *SignupController {
	repo := mymongo.NewUserRepository(app.Mongo, timeout)
	uc := application.SignupUsecase{SignupRepo: repo}
	return &SignupController{&uc}
}

func (lc *SignupController) Signup(
	nickName string,
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*domain.JwtTokenResponse, error) {
	signup, err := domain.NewSignup(
		nickName,
		email,
		password,
		accessTokenSecret,
		accessTokenExpiryHour,
		refreshTokenSecret,
		refreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}
	signupResponse, err := lc.uc.Signup(signup)
	if err != nil {
		return nil, err
	}

	return signupResponse, nil
}
