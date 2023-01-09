package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(name, Id string) (string, error) {
	privateKey, err := ioutil.ReadFile("private.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaprivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"name": name,
		"ID":   Id,
		"exp":  time.Now().Add(30 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(rsaprivateKey)
	if err != nil {
		return "", errors.New("failed to create Token")
	}
	return tokenString, nil
}

func GenerateRefreshToken(name string, Id string) (string, error) {
	privateKey, err := ioutil.ReadFile("private.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaprivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"name": name,
		"ID":   Id,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	refreshTokenString, err := token.SignedString(rsaprivateKey)
	if err != nil {
		return "", errors.New("failed to create Token")
	}
	return refreshTokenString, nil
}
func ValidateJwt(signedToken string) string {
	// Load the RSA public key from a file
	publicKey, err := ioutil.ReadFile("public.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the signed JWT and verify it with the RSA public key
	token, err := jwt.ParseWithClaims(signedToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return rsaPublicKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if token.Valid {
		return "TOKEN IS VALID"
	} else {
		return "TOKEN IS INVALID"
	}
}