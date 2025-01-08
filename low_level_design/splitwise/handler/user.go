package handler

import (
	"encoding/json"
	"net/http"
	constants "splitwise"
	"splitwise/service"
	"splitwise/util"
)

type User struct {
	userService *service.User
	util        *util.Util
}

func NewUserHandler(userService *service.User, util *util.Util) *User {
	return &User{userService: userService, util: util}
}

func (u *User) Me(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie(constants.AUTH_COOKIE)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenClaims, err := u.util.GetJwtClaims(token.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo := u.userService.GetUserByEmail(tokenClaims["email"].(string))
	encodedUser, err := json.Marshal(userInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encodedUser)
}
