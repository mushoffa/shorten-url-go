package main

import (
	"time"

	"shorten-url-go/data/controller"
	"shorten-url-go/data/repository"
	"shorten-url-go/domain/usecase"
	"shorten-url-go/infrastructure/memory"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	repository := repository.NewRepository(memory.NewURLMemory())
	usecase := usecase.NewUsecase(repository)
	controller := controller.NewURLController(usecase)

	server := fiber.New()

	server.Use(limiter.New(limiter.Config{
		Expiration: 5 * time.Second,
		Max: 5,
	}))

	controller.Router(server)

	server.Listen(":9090")
}