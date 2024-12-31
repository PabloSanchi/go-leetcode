package service

import (
	"fmt"
	"log/slog"
	"splitwise/domain/dto"
	"splitwise/repository"
)

const (
	AUTH_ERROR string = "could not login"
)

type Auth struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *Auth {
	return &Auth{userRepository: userRepository}
}

func (as *Auth) Signup(user *dto.User) (*dto.UserInfo, error) {
	var userInfo *dto.UserInfo
	id, err := as.userRepository.CreateUser(user)
	if err != nil {
		slog.Error("error creating user", "error", err)
		return userInfo, fmt.Errorf(AUTH_ERROR)
	}

	userInfo = &dto.UserInfo{
		Id:    id,
		Name:  user.Name,
		Email: user.Email,
	}

	return userInfo, nil
}

func (as *Auth) Login(user *dto.User) (*dto.UserInfo, error) {
	if !as.userRepository.VerifyUser(user) {
		slog.Error("error verifying user")
		var userInfo *dto.UserInfo
		return userInfo, fmt.Errorf(AUTH_ERROR)
	}

	return as.userRepository.GetUserByEmail(user.Email), nil
}

func (as *Auth) VerifyUser(user *dto.User) bool {
	return as.userRepository.VerifyUser(user)
}
