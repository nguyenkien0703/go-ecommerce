package auth

import (
	"example.com/go-ecommerce-backend-api/global"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.Jwt.API_SECRET))
}

func CreateToken(uuidToken string) (string, error) {
	//1. Set time expiration
	timeEx := global.Config.Jwt.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "1h"
	}
	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expiresAt := now.Add(expiration)
	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "shopdevgo",
			Subject:   uuidToken,
		},
	})
}

func ParseJwtTokenSubject(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.API_SECRET), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}

// validate token

func VerifyTokenSubject(token string) (*jwt.StandardClaims, error) {
	claims, err := ParseJwtTokenSubject(token)
	if err != nil {

		return nil, err
	}
	if err = claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}
