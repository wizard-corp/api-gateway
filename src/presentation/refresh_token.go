package presentation

import (
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

func (lc *RefreshTokenController) RefreshToken(
	refreshToken string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*domain.JwtTokenResponse, error) {
	rtc, err := domain.NewRefreshToken(
		refreshToken,
		accessTokenSecret,
		accessTokenExpiryHour,
		refreshTokenSecret,
		refreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}
	rtcResponse, err := lc.uc.RefreshToken(rtc)
	if err != nil {
		return nil, err
	}

	return rtcResponse, nil
}
