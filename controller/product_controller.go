package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	// app.Post("/order-service/seller/login", controller.Login)
	app.Post("/product-service/seller/register", controller.Register)
	app.Post("/product-service/seller/update", controller.Update)
}

func (controller *ProductController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *ProductController) Register(c *fiber.Ctx) error {
	var request model.CreateProduct
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	err = controller.ProductService.Register(request)
	message := "create product successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *ProductController) Update(c *fiber.Ctx) error {
	var request model.UpdateProduct
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	token := c.Get("x-auth-token")

	err = controller.ProductService.UpdateProduct(request, token)
	message := "update product successfull"
	if err != nil {
		message = err.Error()
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}
