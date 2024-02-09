package response

import (
	"authentication-and-authorization-service/internal/domain/constant"
)

var (
	DefaultSuccessResponse = General{
		HTTPStatus: &constant.HTTPStatus200,
		Code:       constant.OK,
		Message:    "success",
	}
)
