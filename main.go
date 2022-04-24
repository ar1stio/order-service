package main

import (
	"order-service/config"
	"order-service/controller"
	"order-service/exception"
	"order-service/repository"
	"order-service/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	// database := config.NewMongoDatabase(configuration)
	sqlDb := config.NewMySqlDatabase(configuration)

	// Setup Repository
	orderRepository := repository.NewOrderRepository(sqlDb)
	buyerRepository := repository.NewBuyerRepository(sqlDb)
	sellerRepository := repository.NewSellerRepository(sqlDb)
	productRepository := repository.NewProductRepository(sqlDb)

	// Setup Service
	orderService := service.NewOrderService(&orderRepository)
	buyerService := service.NewBuyerService(&buyerRepository)
	sellerService := service.NewSellerService(&sellerRepository)
	productService := service.NewProductService(&productRepository)

	// Setup Controller
	orderController := controller.NewOrderController(&orderService)
	buyerController := controller.NewBuyerController(&buyerService)
	sellerController := controller.NewSellerController(&sellerService)
	productController := controller.NewProductController(&productService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Setup Routing
	orderController.Route(app)
	buyerController.Route(app)
	sellerController.Route(app)
	productController.Route(app)

	// Start App
	err := app.Listen(":" + os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
