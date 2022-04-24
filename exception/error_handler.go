package exception

import (
	"net/http"
	"order-service/model"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Error(err.Error())
	status, _ := err.(AuthorizedError)
	if status.Status == "UNAUTHORIZED" {
		return ctx.Status(http.StatusUnauthorized).JSON(model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		})
	}

	statusConflict, _ := err.(ConflictError)
	if statusConflict.Status == "CONFLICT" {
		return ctx.Status(http.StatusConflict).JSON(model.WebResponse{
			Code:   http.StatusConflict,
			Status: "CONFLICT",
			Data:   statusConflict.Message,
		})
	}

	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})

}
