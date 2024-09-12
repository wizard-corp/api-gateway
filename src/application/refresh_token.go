package application

import (
	"errors"

	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/jwttoken"
)

type RefreshTokenUsecase struct {
	RefreshTokenRepo domain.RefreshTokenRepository
}

func (lu *RefreshTokenUsecase) RefreshToken(rtc *domain.RefreshToken) (*domain.JwtTokenResponse, error) {
	id, err := jwttoken.ExtractIDFromToken(rtc.RefreshToken, rtc.JwtToken.RefreshTokenSecret)
	if err != nil {
		return nil, errors.New(domain.NOT_FOUND)
	}

	user, err := lu.RefreshTokenRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New(domain.NOT_FOUND)
	}

	accessToken, err := jwttoken.CreateAccessToken(&user, rtc.JwtToken.AccessTokenSecret, rtc.JwtToken.AccessTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwttoken.CreateRefreshToken(&user, rtc.JwtToken.RefreshTokenSecret, rtc.JwtToken.RefreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	rtcResponse := &domain.JwtTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return rtcResponse, nil
}
