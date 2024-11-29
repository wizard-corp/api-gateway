package application

import (
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/jwttoken"
)

type RefreshTokenUsecase struct {
	RefreshTokenRepo domain.RefreshTokenRepository
}

func (lu *RefreshTokenUsecase) RefreshToken(rtc *domain.RefreshToken) (*domain.JwtTokenResponse, error) {
	id, err := jwttoken.ExtractIDFromToken(rtc.RefreshToken, rtc.JwtToken.RefreshTokenSecret)
	if err != nil {
		return nil, domain.NewInfrastructureError(domain.TOKEN_FAIL, err.Error())
	}

	user, err := lu.RefreshTokenRepo.GetUserByID(id)
	if err != nil {
		return nil, domain.NewApplicationError(domain.DB_FETCH_BY_PARAM, err.Error())
	}

	accessToken, err := jwttoken.CreateAccessToken(&user, rtc.JwtToken.AccessTokenSecret, rtc.JwtToken.AccessTokenExpiryHour)
	if err != nil {
		return nil, domain.NewApplicationError(domain.TOKEN_FAIL, err.Error())
	}

	refreshToken, err := jwttoken.CreateRefreshToken(&user, rtc.JwtToken.RefreshTokenSecret, rtc.JwtToken.RefreshTokenExpiryHour)
	if err != nil {
		return nil, domain.NewApplicationError(domain.TOKEN_FAIL, err.Error())
	}

	rtcResponse := &domain.JwtTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return rtcResponse, nil
}
