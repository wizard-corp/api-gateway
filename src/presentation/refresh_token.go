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

type RefreshTokenController struct {
	uc *application.RefreshTokenUsecase
}

func NewRefreshTokenController(timeout time.Duration, app *bootstrap.App) *RefreshTokenController {
	repo := mymongo.NewUserRepository(app.Mongo, timeout)
	uc := application.RefreshTokenUsecase{RefreshTokenRepo: repo}
	return &RefreshTokenController{&uc}
}

func (lc *RefreshTokenController) NewRefreshToken(
	refreshToken string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*domain.JwtTokenResponse, error) {
	rtc := domain.RefreshToken{
		RefreshToken: refreshToken,
		JwtToken: domain.JwtToken{
			AccessTokenSecret:      accessTokenSecret,
			AccessTokenExpiryHour:  accessTokenExpiryHour,
			RefreshTokenSecret:     refreshTokenSecret,
			RefreshTokenExpiryHour: refreshTokenExpiryHour}}
	errs := rtc.IsRefreshTokenValid()
	if len(errs) > 0 {
		return nil, errors.New(domain.INVALID_SCHEMA + "\n" + strings.Join(errs, "\n"))
	}

	rtcResponse, err := lc.uc.RefreshToken(&rtc)
	if err != nil {
		return nil, err
	}

	return rtcResponse, nil
}
