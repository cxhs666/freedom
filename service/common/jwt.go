package common

import (
	"time"

	"github.com/kataras/jwt"
)

type JwtCreator struct{}

var sharedKey = []byte("sercrethatmaycontainch@r$32chars")

var Jwt = new(JwtCreator)

func (j *JwtCreator) Create(auth Auth) (string, error) {

	token, err := jwt.Sign(jwt.HS256, sharedKey, auth, jwt.MaxAge(24*time.Hour))
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (j *JwtCreator) Verify(token string) (Auth, error) {
	var auth Auth

	tokenByte := []byte(token)
	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, tokenByte)
	if err != nil {
		return auth, err
	}
	err = verifiedToken.Claims(&auth)
	if err != nil {
		return auth, err
	}

	return auth, nil
}
