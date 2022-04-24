package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
)

type BuyerController struct {
	BuyerService service.BuyerService
}

func NewBuyerController(userService *service.BuyerService) BuyerController {
	return BuyerController{BuyerService: *userService}
}

func (controller *BuyerController) Route(app *fiber.App) {
	app.Post("/order-service/buyer/login", controller.Login)
	app.Post("/order-service/buyer/register", controller.Register)
	app.Post("/order-service/buyer/update", controller.UpdateBuyer)
	// app.Post("/order-service/user/authentication", controller.AuthenticationBuyer)
	// app.Post("/order-service/user/register-member", controller.RegisterMember)
}

func (controller *BuyerController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *BuyerController) Login(c *fiber.Ctx) error {
	var request model.LoginBuyerReq
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.BuyerService.Login(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *BuyerController) Register(c *fiber.Ctx) error {
	var request model.CreateBuyer
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	err = controller.BuyerService.Register(request)
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

func (controller *BuyerController) UpdateBuyer(c *fiber.Ctx) error {
	var request model.UpdateBuyer
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.BuyerService.UpdateBuyer(request, token)
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
