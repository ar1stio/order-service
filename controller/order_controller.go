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

func (controller *OrderController) RouteOrder(app *fiber.App) {

	app.Post("/order-service/show-order-list", controller.ShowOrderList)
	app.Post("/order-service/show-order-product", controller.ShowOrderProduct)

	app.Post("/order-service/buyer/order", controller.CreateOrder)
	app.Post("/order-service/delivered", controller.Delivered)
}

func (controller *OrderController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *OrderController) ShowOrderList(c *fiber.Ctx) error {
	var request model.GetFindOrder
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	data := controller.OrderService.FindsOrderList(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *OrderController) ShowOrderProduct(c *fiber.Ctx) error {
	var request model.GetFindOrder
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	data := controller.OrderService.FindsOrderProduct(request)

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
	// token := c.Get("x-auth-token")

	err = controller.OrderService.Register(request)
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

func (controller *OrderController) Delivered(c *fiber.Ctx) error {
	var request model.Deliveredreq
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	// token := c.Get("x-auth-token")

	err = controller.OrderService.Delrivered(request)
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
