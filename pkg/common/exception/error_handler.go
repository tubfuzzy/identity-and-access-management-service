package exception

import (
	"authentication-and-authorization-service/internal/domain/constant"

	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var resp *ErrorResponse
	switch err.(type) {
	case ValidationError, BadRequestError:
		resp = ErrBadRequest
		resp.Errors = err.Error()
	case NotFoundError:
		resp = ErrNotFound
		resp.Errors = err.Error()
	case UnauthorizedError:
		resp = ErrUnauthenticated
	default:
		resp = &DefaultErrorResponse
	}
	var e *fiber.Error
	if errors.As(err, &e) {
		resp.Code = constant.FRAMEWORK_ERROR
		resp.Message = e.Message
	}

	return ctx.Status(*resp.HTTPStatus).JSON(resp)
}
