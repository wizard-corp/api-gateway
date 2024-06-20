package presentation

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/wizard-corp/api-gateway/domain"
)

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewLoginController(
	uc domain.LoginUsecase,
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int,
) (*LoginResponse, error) {
	user, err := uc.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New(domain.USER_NOT_FOUND)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New(domain.INVALID_CREDENTIALS)
	}

	accessToken, err := uc.CreateAccessToken(&user, accessTokenSecret, accessTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.CreateRefreshToken(&user, refreshTokenSecret, refreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
