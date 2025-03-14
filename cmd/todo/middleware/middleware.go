package middleware

import (
	"errors"
	"todo/internal/serrors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ErrorResponseHandler(ctx *fiber.Ctx) error {
	var serviceError serrors.ServiceError
	err := ctx.Next()
	if err != nil {
		log.Error(err)
		if errors.As(err, &serviceError) {
			ctx.Status(serviceError.GetHttpCode())
			return ctx.SendString(err.Error())
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return nil
}
