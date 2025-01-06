package util

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"splitwise"
	"splitwise/domain/dto"
	"time"
)

type Util struct{}

func NewUtil() *Util {
	return &Util{}
}

func (u *Util) GenerateJwt(user *dto.UserInfo) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":  user.Name,
			"email": user.Email,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	tokenString, err := token.SignedString(constants.SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *Util) ValidateJwt(tokenString string) (bool, error) {
	token, err := parseJwt(tokenString)
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (u *Util) GetJwtClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := parseJwt(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func (u *Util) HashPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func (u *Util) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func parseJwt(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return constants.SECRET_KEY, nil
	})
}
