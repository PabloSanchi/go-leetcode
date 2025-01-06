package main

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
	constants "splitwise"
	"splitwise/config"
	"splitwise/handler"
	"splitwise/middleware"
	"splitwise/repository"
	"splitwise/service"
	"splitwise/util"
)

func main() {
	db, err := config.InitDb()
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		os.Exit(1)
	}

	slog.Info("connection to db successful")
	slog.Info("server started successfully", "port", constants.PORT)

	utils := util.NewUtil()
	md := middleware.NewMiddleware(utils)

	userRepository := repository.NewUserRepositoryImpl(db, utils)
	groupRepository := repository.NewGroupRepositoryImpl(db)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	groupService := service.NewGroupService(groupRepository)
	authHandler := handler.NewAuthHandler(authService, utils)
	userHandler := handler.NewUserHandler(userService, utils)
	groupHandler := handler.NewGroupHandler(groupService)

	router := mux.NewRouter()

	router.HandleFunc(endpoint("auth/signup"), authHandler.Signup)
	router.HandleFunc(endpoint("auth/login"), authHandler.Login)
	router.HandleFunc(endpoint("auth/logout"), authHandler.Logout)

	router.HandleFunc(endpoint("users/me"), md.WithAuth(userHandler.Me))

	router.HandleFunc(endpoint("groups"), md.WithAuth(groupHandler.CreateGroup))
	router.HandleFunc(endpoint("groups/{id}"), md.WithAuth(groupHandler.GetGroup))
	router.HandleFunc(endpoint("groups/{id}/users"), md.WithAuth(groupHandler.AddUsersToGroup))

	if err := http.ListenAndServe(constants.PORT, router); err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(1)
	}

}

func endpoint(path string) string {
	return "/api/" + constants.VERSION + "/" + path
}
