package main

import (
	"log/slog"
	"net/http"
	"os"
	"splitwise/config"
	"splitwise/handler"
	"splitwise/middleware"
	"splitwise/repository"
	"splitwise/service"
)

const (
	VERSION string = "v1"
	PORT    string = ":8080"
)

func main() {
	db, err := config.InitDb()
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		os.Exit(1)
	}

	slog.Info("connection to db successful")
	slog.Info("server started successfully", "port", PORT)

	middleware := middleware.NewMiddleware()

	userRepository := repository.NewUserRepositoryImpl(db)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(authService)
	userhandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc(endpoint("auth/signup"), authHandler.Signup)
	mux.HandleFunc(endpoint("auth/login"), authHandler.Login)
	mux.HandleFunc(endpoint("auth/logout"), authHandler.Logout)

	mux.HandleFunc(endpoint("users/me"), middleware.WithAuth(userhandler.Me))

	if err := http.ListenAndServe(PORT, mux); err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(1)
	}

}

func endpoint(path string) string {
	return "/api/" + VERSION + "/" + path
}
