package controller

import (
	"order-service/exception"
	"order-service/model"
	"order-service/service"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/product-service/seller-product/register", controller.Register)
	app.Post("/product-service/seller-product/update", controller.Update)
	app.Get("/product-service/product/:id", controller.FindSingleProduct)
	app.Post("/product-service/all-product", controller.FindAllProduct)
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

func (controller *ProductController) FindSingleProduct(c *fiber.Ctx) error {
	// err := c.BodyParser(&request)
	// exception.PanicIfNeeded(err)
	id := c.Params("id")
	idVar, _ := strconv.Atoi(id)

	data := controller.ProductService.FindSingleProduct(idVar)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

func (controller *ProductController) FindAllProduct(c *fiber.Ctx) error {
	var request model.GetAllProductReq
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	data := controller.ProductService.FindAllProduct(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}
