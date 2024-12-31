package service

import (
	"fmt"
	"log/slog"
	"splitwise/domain/dto"
	"splitwise/repository"
	"splitwise/util"
)

const (
	AUTH_ERROR string = "could not login"
)

type Auth struct {
	userRepository repository.UserRepository
	util           *util.Util
}

func NewAuthService(userRepository repository.UserRepository) *Auth {
	return &Auth{
		userRepository: userRepository,
		util:           util.NewUtil(),
	}
}

func (as *Auth) Signup(user *dto.User) (string, error) {
	if _, err := as.userRepository.CreateUser(user); err != nil {
		slog.Error("error creating user", "error", err)
		return "", fmt.Errorf(AUTH_ERROR)
	}

	token, err := as.util.GenerateJwt(user.Email)
	if err != nil {
		slog.Error("error generating token", "error", err)
		return "", fmt.Errorf(AUTH_ERROR)
	}
	return token, nil
}

func (as *Auth) Login(user *dto.User) (string, error) {
	if !as.userRepository.VerifyUser(user) {
		slog.Error("error verifying user")
		return "", fmt.Errorf(AUTH_ERROR)
	}

	token, err := as.util.GenerateJwt(user.Email)
	if err != nil {
		slog.Error("error generating token", "error", err)
		return "", fmt.Errorf(AUTH_ERROR)
	}

	return token, nil
}

func (as *Auth) VerifyUser(user *dto.User) bool {
	return as.userRepository.VerifyUser(user)
}
