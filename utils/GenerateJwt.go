package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

var Wrapper JwtWrapper

func init() {
	if err := godotenv.Load(); err != nil {
		println("No .env file found")
	}

	expireTime, errParsing := strconv.ParseInt(os.Getenv("JWT_EXPIRES"), 10, 64)
	if errParsing != nil {
		log.Fatal("Error parsing JWT_EXPIRES")
	}

	Wrapper = JwtWrapper{
		SecretKey:       os.Getenv("JWT_SECRET"),
		Issuer:          os.Getenv("JWT_ISSUER"),
		ExpirationHours: expireTime,
	}
}

func GenerateToken(email string) (signedToken string, err error) {
	claims := JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(Wrapper.ExpirationHours)).Unix(),
			Issuer:    Wrapper.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(Wrapper.SecretKey))
	if err != nil {
		return
	}

	return signedToken, nil
}

func VerifyToken(token string) (claims *JwtClaim, err error) {
	claims = &JwtClaim{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Wrapper.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	return claims, nil
}
