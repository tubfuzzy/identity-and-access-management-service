package response

import (
	"github.com/gofiber/fiber/v2"
)

type General struct {
	HTTPStatus *int        `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func (g *General) JSON(c *fiber.Ctx) error {
	return c.Status(*g.HTTPStatus).JSON(g)
}
