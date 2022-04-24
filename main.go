package main

import (
	"os"
	"order-service/config"
	"order-service/controller"
	"order-service/exception"
	"order-service/repository"
	"order-service/service"

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

	// Setup Service
	orderService := service.NewOrderService(&orderRepository)

	// Setup Controller
	orderController := controller.NewOrderController(&orderService)
	// addressController := controller.NewAddressController(&orderService, &addressService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Setup Routing
	orderController.Route(app)

	// Start App
	err := app.Listen(":" + os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
