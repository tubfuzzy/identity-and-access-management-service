package http

import (
	"authentication-and-authorization-service/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler struct {
	domain.AuthenticationService
}

func NewAuthenticationHandler(authenticationService domain.AuthenticationService) AuthenticationHandler {
	return AuthenticationHandler{
		AuthenticationService: authenticationService,
	}
}

func (h AuthenticationHandler) InitRoute(app fiber.Router) {
	app.Get("/test", h.TestHandler)
}

func (h AuthenticationHandler) TestHandler(ctx *fiber.Ctx) error {
	response := map[string]interface{}{
		"message": "This is a sample response for the /test route",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
