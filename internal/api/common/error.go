package common

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type BusinessError struct {
	Code    Result
	Message string
}

func (e *BusinessError) Error() string {
	return e.Message
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var resp *APIResponse

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		resp = MakeAPIResponseWithMsg(ResultFailed, nil, fiberErr.Message)
		return c.Status(fiberErr.Code).JSON(resp)
	}

	var businessErr *BusinessError
	if errors.As(err, &businessErr) {
		if len(businessErr.Message) > 0 {
			resp = MakeAPIResponseWithMsg(businessErr.Code, nil, businessErr.Message)
		} else {
			resp = MakeAPIResponse(businessErr.Code, nil)
		}
		return c.Status(businessErr.Code.HttpCode).JSON(resp)
	}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		result := ResultFormInvalid
		resp = MakeAPIResponse(result, validationErrs.Error())
		return c.Status(result.HttpCode).JSON(resp)
	}

	zap.S().Errorw(
		"Unknown error handled",
		"err", err,
	)
	unknown := ResultUnknown
	return c.Status(unknown.HttpCode).
		JSON(MakeAPIResponseWithMsg(unknown, nil, err.Error()))
}
