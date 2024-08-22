package domain

import (
	"regexp"

	"github.com/wizard-corp/api-gateway/mymongo"
)

type Login struct {
	Email                  string
	Password               string
	AccessTokenSecret      string
	AccessTokenExpiryHour  int
	RefreshTokenSecret     string
	RefreshTokenExpiryHour int
}

type LoginUsecase interface {
	GetUserByEmail(email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type loginUsecase struct {
	writeUId       string
	userRepository UserRepository
}

func NewLoginUsecase(writeUId string, db mymongo.MongoServer) *loginUsecase {
	return &loginUsecase{
		writeUId:       writeUId,
		userRepository: mymongo.NewUserRepository(db, "user"),
	}
}

func IsLoginValid(email string, password string) []string {
	var errors []string

	if email == "" {
		errors = append(errors, "email is empty")
	}

	if password == "" {
		errors = append(errors, "password is empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		errors = append(errors, "invalid format for email")
	}

	return errors
}
