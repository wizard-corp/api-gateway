package application

import (
	"github.com/wizard-corp/api-gateway/src/domain"
	"github.com/wizard-corp/api-gateway/src/jwttoken"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	LoginRepo domain.LoginRepository
}

func (lu *LoginUsecase) Login(login *domain.Login) (*domain.JwtTokenResponse, error) {
	user, err := lu.LoginRepo.GetUserByEmail(login.Email)
	if err != nil {
		return nil, domain.NewInfrastructureError(domain.DB_FETCH_BY_PARAM, err.Error())
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)) != nil {
		return nil, domain.NewApplicationError(domain.COMPARE_FAIL, "Incorrect password")
	}

	accessToken, err := jwttoken.CreateAccessToken(&user, login.JwtToken.AccessTokenSecret, login.JwtToken.AccessTokenExpiryHour)
	if err != nil {
		return nil, domain.NewApplicationError(domain.TOKEN_FAIL, err.Error())
	}

	refreshToken, err := jwttoken.CreateRefreshToken(&user, login.JwtToken.RefreshTokenSecret, login.JwtToken.RefreshTokenExpiryHour)
	if err != nil {
		return nil, domain.NewApplicationError(domain.TOKEN_FAIL, err.Error())
	}

	loginResponse := &domain.JwtTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return loginResponse, nil
}
