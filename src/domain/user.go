package domain

import (
	"net/mail"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	NickName string             `bson:"nickName"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	PersonID primitive.ObjectID `bson:"personId"`
}

type UserRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}

func (u *User) IsValidUser() []string {
	var errs []string

	if u.Email == "" {
		errs = append(errs, IS_EMPTY)
	}

	if u.Password == "" {
		errs = append(errs, IS_EMPTY)
	}

	if IsValidEmail(u.Email) {
		errs = append(errs, INVALID_FORMAT)
	}

	password_errs := IsValidPassword(u.Password)
	if len(password_errs) > 0 {
		errs = append(errs, password_errs...)
	}

	return errs
}

func (u *User) MongoSanitizeUser() {
	u.Email = strings.Trim(u.Email, " $/^\\")
	u.Password = strings.Trim(u.Password, " $/^\\")
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func IsValidPassword(password string) []string {
	var errs []string

	if len(password) < 8 {
		errs = append(errs, "Password should be of 8 characters long")
	}
	done, _ := regexp.MatchString("([a-z])+", password)
	if !done {
		errs = append(errs, "Password should contain atleast one lower case character")
	}
	done, _ = regexp.MatchString("([A-Z])+", password)
	if !done {
		errs = append(errs, "Password should contain atleast one upper case character")
	}
	done, _ = regexp.MatchString("([0-9])+", password)
	if !done {
		errs = append(errs, "Password should contain atleast one digit")
	}
	done, _ = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if !done {
		errs = append(errs, "Password should contain atleast one special character")
	}

	return errs
}
