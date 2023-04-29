package helper

import (
	"belajar-golang-gql/models/web"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// compare login password with hash password from db
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(request web.TokenCreateRequest, value time.Duration) (string, error) {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	var APPLICATION_NAME = "BELAJAR"

	expiredTime := time.Now().Add(time.Hour * value).Unix()

	claims := &web.TokenClaims{
		UserID:    request.UserID,
		Email:     request.Email,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expiredTime,
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := tokens.SignedString(jwtTokenSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func TokenClaims(userToken string) (*web.TokenClaims, error) {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &web.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtTokenSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil

}
