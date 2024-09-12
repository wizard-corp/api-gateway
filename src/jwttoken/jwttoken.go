package jwttoken

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/wizard-corp/api-gateway/src/domain"
)

const (
	INVALID_TOKEN    = "Invalid token\n"
	INCORRECT_METHOD = "Unexpected singing method\n"
	DECODE_FAIL      = "failed to parse\n"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	claims := &JwtCustomClaims{
		Name: user.NickName,
		ID:   user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			txt := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return false, errors.New(DECODE_FAIL + "\n" + txt)
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, errors.New(DECODE_FAIL + "\n" + err.Error())
	}

	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			txt := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(DECODE_FAIL + "\n" + txt)
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", errors.New(DECODE_FAIL + "\n" + err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", errors.New(INVALID_TOKEN)
	}

	return claims["id"].(string), nil
}
