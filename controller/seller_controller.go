package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
)

type SellerRegisterController struct {
	SellerService service.SellerService
}

func NewSellerRegisterController(userService *service.SellerService) SellerRegisterController {
	return SellerRegisterController{SellerService: *userService}
}

func (controller *SellerRegisterController) Route(app *fiber.App) {
	app.Post("/order-service/seller/login", controller.Login)
	app.Post("/order-service/seller/register", controller.Register)
}

func (controller *SellerRegisterController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *SellerRegisterController) Login(c *fiber.Ctx) error {
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

func (controller *SellerRegisterController) Register(c *fiber.Ctx) error {
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
