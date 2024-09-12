package application

import (
	"errors"

	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/jwttoken"
	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	SignupRepo domain.SignupRepository
}

func (lu *SignupUsecase) Signup(signup *domain.Signup) (*domain.JwtTokenResponse, error) {
	_, err := lu.SignupRepo.GetUserByEmail(signup.User.Email)
	if err == nil {
		return nil, errors.New(domain.ALREADY_EXISTS)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(signup.User.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, errors.New(domain.ENCRIPTION_ERROR)
	}
	signup.User.Password = string(encryptedPassword)

	err = lu.SignupRepo.Create(&signup.User)
	if err != nil {
		return nil, errors.New(domain.DB_INSERT_FAILED)
	}

	accessToken, err := jwttoken.CreateAccessToken(&signup.User, signup.JwtToken.AccessTokenSecret, signup.JwtToken.AccessTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwttoken.CreateRefreshToken(&signup.User, signup.JwtToken.RefreshTokenSecret, signup.JwtToken.RefreshTokenExpiryHour)
	if err != nil {
		return nil, err
	}

	signupResponse := &domain.JwtTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return signupResponse, nil
}