package application

import (
	"errors"

	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/jwttoken"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	LoginRepo domain.LoginRepository
}

func (lu *LoginUsecase) NewLogin(login *domain.Login) (*domain.JwtTokenResponse, error) {
	user, err := lu.LoginRepo.GetUserByEmail(login.Email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)) != nil {
		return nil, errors.New(domain.INVALID_CREDENTIALS)
	}

	accessToken, err := jwttoken.CreateAccessToken(&user, login.JwtToken.AccessTokenSecret, login.JwtToken.AccessTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwttoken.CreateRefreshToken(&user, login.JwtToken.RefreshTokenSecret, login.JwtToken.RefreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	loginResponse := &domain.JwtTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return loginResponse, nil
}
