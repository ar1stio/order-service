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
	buyerController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

// Setup Repository
var sqlDb = config.NewMySqlDatabase(configuration)

var buyerRepository = repository.NewBuyerRepository(sqlDb)

// Setup Service
var buyerService = service.NewBuyerService(&buyerRepository)

// Setup Controller
var buyerController = NewBuyerController(&buyerService)

var app = createTestApp()
