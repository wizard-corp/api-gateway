package application

import (
	"errors"

	"github.com/wizard-corp/api-gateway/domain"
	"github.com/wizard-corp/api-gateway/jwttoken"
	"golang.org/x/crypto/bcrypt"
)

func (uc *domain.LoginUsecase) NewLogin(
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int,
) (domain.LoginResponse, error) {
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New(domain.INVALID_CREDENTIALS)
	}

	accessToken, err := jwttoken.CreateAccessToken(user, accessTokenSecret, accessTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwttoken.CreateRefreshToken(user, refreshTokenSecret, refreshTokenExpiryHour)
	if err != nil {
		return err
	}

	loginResponse := &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return loginResponse, nil
}
