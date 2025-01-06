package middleware

import (
	"net/http"
	"splitwise"
	"splitwise/util"
)

type Middleware struct {
	util *util.Util
}

func NewMiddleware(util *util.Util) *Middleware {
	return &Middleware{util: util}
}

func (m *Middleware) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie(constants.AUTH_COOKIE)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ok, err := m.util.ValidateJwt(token.Value)
		if err != nil || !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
