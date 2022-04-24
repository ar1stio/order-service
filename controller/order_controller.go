package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	OrderService service.OrderService
}

func NewOrderController(orderService *service.OrderService) OrderController {
	return OrderController{OrderService: *orderService}
}

func (controller *OrderController) Route(app *fiber.App) {

	app.Post("/order-service/company/show-member", controller.ShowMember)

	app.Post("/order-service/buyer/order", controller.CreateOrder)
	app.Post("/order-service/admin/update-order", controller.UpdateOrder)
	app.Post("/order-service/admin/activate-order", controller.ActivateOrder)
	app.Post("/order-service/admin/nonactivate-order", controller.NonActivateOrder)
	app.Post("/order-service/admin/get-token-order", controller.FindDataOrder)
}

func (controller *OrderController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *OrderController) ShowMember(c *fiber.Ctx) error {
	var request model.GetOrderCompanyRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	data := controller.OrderService.FinsOrderCompany(request.MemberCompanyId)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *OrderController) CreateOrder(c *fiber.Ctx) error {
	var request model.CreateOrder
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.OrderService.CreateOrder(request, token)
	message := "create order successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *OrderController) UpdateOrder(c *fiber.Ctx) error {
	var request model.UpdateOrder
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.OrderService.UpdateOrder(request, token)
	message := "update order successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *OrderController) ActivateOrder(c *fiber.Ctx) error {
	var request model.UpdateActivateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.OrderService.ActivateOrder(request, token)
	message := "update activate successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *OrderController) NonActivateOrder(c *fiber.Ctx) error {
	var request model.NonActivateOrder
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.OrderService.NonActivateOrder(request, token)
	message := "update nonactivate successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *OrderController) FindDataOrder(c *fiber.Ctx) error {
	var request model.GetLoginRequest
	err := c.BodyParser(&request)
	token := c.Get("x-auth-token")

	exception.PanicIfNeeded(err)

	response := controller.OrderService.FindTokenOrder(request, token)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
