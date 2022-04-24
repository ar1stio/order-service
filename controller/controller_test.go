package controller

import (
	"order-service/config"
	"order-service/repository"
	"order-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	orderController.RouteOrder(app)
	return app
}

var configuration = config.New("../.env.test")

// Setup Repository
var sqlDb = config.NewMySqlDatabase(configuration)

var orderRepository = repository.NewOrderRepository(sqlDb)

// Setup Service
var orderService = service.NewOrderService(&orderRepository)

// Setup Controller
var orderController = NewOrderController(&orderService)

var app = createTestApp()
