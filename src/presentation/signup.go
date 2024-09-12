package presentation

import (
	"errors"
	"strings"
	"time"

	"github.com/wizard-corp/api-gateway/src/application"
	"github.com/wizard-corp/api-gateway/src/bootstrap"
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/mymongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupController struct {
	uc *application.SignupUsecase
}

func NewSignupController(timeout time.Duration, app *bootstrap.App) *SignupController {
	repo := mymongo.NewUserRepository(app.Mongo, timeout)
	uc := application.SignupUsecase{SignupRepo: repo}
	return &SignupController{&uc}
}

func (lc *SignupController) NewSignup(
	nickName string,
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*domain.JwtTokenResponse, error) {
	signup := domain.Signup{
		User: domain.User{
			ID:       primitive.NewObjectID(),
			NickName: nickName,
			Email:    email,
			Password: password},
		JwtToken: domain.JwtToken{
			AccessTokenSecret:      accessTokenSecret,
			AccessTokenExpiryHour:  accessTokenExpiryHour,
			RefreshTokenSecret:     refreshTokenSecret,
			RefreshTokenExpiryHour: refreshTokenExpiryHour}}
	errs := signup.IsSignupValid()
	if len(errs) > 0 {
		txt := domain.INVALID_SCHEMA + "\n" + strings.Join(errs, "\n")
		return nil, errors.New(txt)
	}

	signupResponse, err := lc.uc.Signup(&signup)
	if err != nil {
		return nil, err
	}

	return signupResponse, nil
}
