package domain

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signup struct {
	User
	JwtToken
}

type SignupRepository interface {
	GetUserByEmail(email string) (User, error)
	Create(*User) error
}

func (s *Signup) IsSignupValid() []string {
	errors := s.User.IsValidUser()
	return errors
}

func NewSignup(
	nickName string,
	email string,
	password string,
	accessTokenSecret string,
	accessTokenExpiryHour int,
	refreshTokenSecret string,
	refreshTokenExpiryHour int) (*Signup, error) {
	signup := Signup{
		User: User{
			ID:       primitive.NewObjectID(),
			NickName: nickName,
			Email:    email,
			Password: password},
		JwtToken: JwtToken{
			AccessTokenSecret:      accessTokenSecret,
			AccessTokenExpiryHour:  accessTokenExpiryHour,
			RefreshTokenSecret:     refreshTokenSecret,
			RefreshTokenExpiryHour: refreshTokenExpiryHour}}
	errs := signup.IsSignupValid()
	if len(errs) > 0 {
		return nil, NewDomainError(INVALID_SCHEMA, strings.Join(errs, "\n"))
	}
	return &signup, nil
}
