package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
)

type BuyerRegisterController struct {
	BuyerService service.BuyerService
}

func NewBuyerRegisterController(buyerService *service.BuyerService) BuyerRegisterController {
	return BuyerRegisterController{BuyerService: *buyerService}
}

func (controller *BuyerRegisterController) Route(app *fiber.App) {
	app.Post("/order-service/buyer/login", controller.Login)
	app.Post("/order-service/buyer/register", controller.Register)
}

func (controller *BuyerRegisterController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *BuyerRegisterController) Login(c *fiber.Ctx) error {
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

func (controller *BuyerRegisterController) Register(c *fiber.Ctx) error {
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
