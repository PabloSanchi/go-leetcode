package main

import (
	"fmt"
	"log/slog"
)

type Request struct {
	Role     string
	LoggedIn bool
}

func NewRequest(role string, loggedIn bool) *Request {
	return &Request{
		Role:     role,
		LoggedIn: loggedIn,
	}
}

type RequestHandler interface {
	Handle(request *Request) bool
	SetNext(handler RequestHandler)
}

type AuthenticatorHandler struct {
	next RequestHandler
}

func (e *AuthenticatorHandler) Handle(request *Request) bool {
	slog.Info("Authenticator Handler")
	if !request.LoggedIn {
		return false
	}

	if e.next != nil {
		return e.next.Handle(request)
	}

	return true
}

func (e *AuthenticatorHandler) SetNext(handler RequestHandler) {
	e.next = handler
}

type AuthorizationHandler struct {
	next RequestHandler
}

func (o *AuthorizationHandler) Handle(request *Request) bool {
	slog.Info("Authorization Handler")
	if request.Role != "ADMIN" {
		return false
	}

	if o.next != nil {
		return o.next.Handle(request)
	}

	return true
}

func (o *AuthorizationHandler) SetNext(handler RequestHandler) {
	o.next = handler
}

func main() {
	authenticator := &AuthenticatorHandler{}
	authorization := &AuthorizationHandler{}

	authenticator.SetNext(authorization)

	authenticatedAndAuthorizedReq := NewRequest("ADMIN", true)
	fmt.Println(authenticator.Handle(authenticatedAndAuthorizedReq))

	authenticatedOnly := NewRequest("DEFAULT", true)
	fmt.Println(authenticator.Handle(authenticatedOnly))

	notAuthenticatedNorAuthorized := NewRequest("DEFAULT", false)
	fmt.Println(authenticator.Handle(notAuthenticatedNorAuthorized))
}
