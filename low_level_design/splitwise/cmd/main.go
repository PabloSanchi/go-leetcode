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
	"splitwise/util"
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

	utils := util.NewUtil()
	md := middleware.NewMiddleware(utils)

	userRepository := repository.NewUserRepositoryImpl(db, utils)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(authService, utils)
	userHandler := handler.NewUserHandler(userService, utils)

	mux := http.NewServeMux()

	mux.HandleFunc(endpoint("auth/signup"), authHandler.Signup)
	mux.HandleFunc(endpoint("auth/login"), authHandler.Login)
	mux.HandleFunc(endpoint("auth/logout"), authHandler.Logout)

	mux.HandleFunc(endpoint("users/me"), md.WithAuth(userHandler.Me))

	if err := http.ListenAndServe(PORT, mux); err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(1)
	}

}

func endpoint(path string) string {
	return "/api/" + VERSION + "/" + path
}
