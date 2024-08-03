package jwtoken

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	errWrongMethod       = errors.New("wrong sign in method")
	errInvalidTokenValue = errors.New("wrong token value")
)

type Auth struct {
	salt   string
	jwtTTL time.Duration
	jwtKey string
}

func NewAuth(salt string, jwtTTL time.Duration, jwtKey string) *Auth {
	return &Auth{
		salt:   salt,
		jwtTTL: jwtTTL,
		jwtKey: jwtKey,
	}
}

func (a *Auth) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(a.salt)))
}

func (a *Auth) GenerateToken(id int) (signedString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.Itoa(id),
		"exp": time.Now().Add(a.jwtTTL).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(a.jwtKey))
}

func (a *Auth) ParseToken(signedString string) (id int, err error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errWrongMethod
		}

		return []byte(a.jwtKey), nil
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
