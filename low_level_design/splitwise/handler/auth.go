// Package handler Auth | Signup, Login and Logout requests
// For an improved security we always return bad request error, so the attacker can't make guesses about the user
package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"splitwise/domain/dto"
	"splitwise/middleware"
	"splitwise/service"
	"splitwise/util"
)

type Auth struct {
	authService *service.Auth
	util        *util.Util
}

func NewAuthHandler(authService *service.Auth) *Auth {
	return &Auth{authService: authService, util: util.NewUtil()}
}

func (ah *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, err := ah.authService.Signup(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := ah.util.GenerateJwt(userInfo)
	if err != nil {
		slog.Error("error generating token", "error", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	http.SetCookie(w, createAuthCookie(token))
	w.WriteHeader(http.StatusCreated)
}

func (ah *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, err := ah.authService.Login(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwtString, err := ah.util.GenerateJwt(userInfo)
	if err != nil {
		slog.Error("error generating token", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w, createAuthCookie(jwtString))
	w.WriteHeader(http.StatusOK)
}

func (ah *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, createAuthCookie(""))
	w.WriteHeader(http.StatusOK)
}

func createAuthCookie(jwtString string) *http.Cookie {
	return &http.Cookie{
		Name:     middleware.AUTH_COOKIE,
		Value:    jwtString,
		Path:     "/",
		HttpOnly: true,
	}
}
