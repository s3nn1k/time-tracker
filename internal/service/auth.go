package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt = "asdfnvcxb435yt"

	tokenTTL  = 30 * time.Minute
	signInKey = "ieowlWERJilea234eriuyaYUNKnb9283"
)

var (
	errWrongMethod       = errors.New("wrong sign in method")
	errInvalidTokenValue = errors.New("wrong token value")
)

type Auth struct{}

func newAuth() *Auth {
	return &Auth{}
}

func (a *Auth) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *Auth) GenerateToken(id int) (signedString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.Itoa(id),
		"exp": time.Now().Add(tokenTTL).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(signInKey))
}

func (a *Auth) ParseToken(signedString string) (id int, err error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errWrongMethod
		}

		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errInvalidTokenValue
	}

	claims := token.Claims.(jwt.MapClaims)
	strId := claims["id"].(string)

	id, err = strconv.Atoi(strId)
	if err != nil {
		return 0, err
	}

	return id, nil
}
