package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
)

type SellerController struct {
	SellerService service.SellerService
}

func NewSellerController(userService *service.SellerService) SellerController {
	return SellerController{SellerService: *userService}
}

func (controller *SellerController) Route(app *fiber.App) {
	app.Post("/order-service/seller/login", controller.Login)
	app.Post("/order-service/seller/register", controller.Register)
	app.Post("/order-service/seller/update", controller.Update)
}

func (controller *SellerController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *SellerController) Login(c *fiber.Ctx) error {
	var request model.LoginSellerReq
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.SellerService.Login(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *SellerController) Register(c *fiber.Ctx) error {
	var request model.CreateSeller
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	err = controller.SellerService.Register(request)
	message := "create user successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *SellerController) Update(c *fiber.Ctx) error {
	var request model.UpdateSeller
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.SellerService.UpdateSeller(request, token)
	message := "update user successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}
