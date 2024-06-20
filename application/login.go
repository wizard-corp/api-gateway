package application

import (
	"github.com/wizard-corp/api-gateway/domain"
	"github.com/wizard-corp/api-gateway/jwttoken"
	"github.com/wizard-corp/api-gateway/mymongo"
)

type loginUsecase struct {
	writeUId       string
	userRepository domain.UserRepository
}

func NewLoginUsecase(writeUId string, db mymongo.MongoDB) domain.LoginUsecase {
	return &loginUsecase{
		writeUId:       writeUId,
		userRepository: mymongo.NewUserRepository(db, "user"),
	}
}

func (uc *loginUsecase) GetUserByEmail(email string) (domain.User, error) {
	return uc.userRepository.GetByEmail(email)
}

func (uc *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return jwttoken.CreateAccessToken(user, secret, expiry)
}

func (uc *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return jwttoken.CreateRefreshToken(user, secret, expiry)
}
