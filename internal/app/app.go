package app

import (
	authenticationhandler "authentication-and-authorization-service/internal/authentication/controller/http"
	authenticationrepositorydb "authentication-and-authorization-service/internal/authentication/repository/db"
	authenticationservice "authentication-and-authorization-service/internal/authentication/service"

	"authentication-and-authorization-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func NewApplication(app fiber.Router, logger logger.Logger) {
	authenticationRepositorydb := authenticationrepositorydb.NewAuthenticationRepository()
	authenticationService := authenticationservice.NewAuthenticationService(authenticationRepositorydb)
	authenticationHandler := authenticationhandler.NewAuthenticationHandler(authenticationService)

	authenticationHandler.InitRoute(app)
}
