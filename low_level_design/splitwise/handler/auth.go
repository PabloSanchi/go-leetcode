// Package handler Auth | Signup, Login and Logout requests
// For an improved security we always return bad request error, so the attacker can't make guesses about the user
package handler

import (
	"encoding/json"
	"net/http"
	"splitwise/domain/dto"
	"splitwise/middleware"
	"splitwise/service"
)

type Auth struct {
	authService *service.Auth
}

func NewAuthHandler(authService *service.Auth) *Auth {
	return &Auth{authService: authService}
}

func (ah *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwtString, err := ah.authService.Signup(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w, createAuthCookie(jwtString))
	w.WriteHeader(http.StatusCreated)
}

func (ah *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwtString, err := ah.authService.Login(&user)
	if err != nil {
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
