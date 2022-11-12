package controller

import (
	"shorten-url-go/data/model/controller"
	"shorten-url-go/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type URLController interface {
	Router(*fiber.App)
	Encode(*fiber.Ctx) error
	Redirect(*fiber.Ctx) error
	FindByUrl(*fiber.Ctx) error
	FindAllUrl(*fiber.Ctx) error
}

type url struct {
	usecase usecase.Usecase
}

func NewURLController(usecase usecase.Usecase) URLController {
	return &url{usecase}
}

func (c *url) Router(r *fiber.App) {
	r.Post("/encode", c.Encode)
	r.Get("/r/:url", c.Redirect)
	r.Get("/findByUrl/:url", c.FindByUrl)
	r.Get("/findAll", c.FindAllUrl)
}

func (c *url) Encode(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	request := model.EncodeRequest{}
	response := model.BaseResponse{}

	if err := ctx.BodyParser(&request); err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	data, err := c.usecase.EncodeURL(request.URL)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{ "shorten_url": data })
}

func (c *url) Redirect(ctx *fiber.Ctx) error {
	response := model.BaseResponse{}

	shortenURL := ctx.Params("url")

	originalURL, err := c.usecase.DecodeURL(shortenURL)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return ctx.Redirect(originalURL, fiber.StatusTemporaryRedirect)
}

func (c *url) FindByUrl(ctx *fiber.Ctx) error {
	response := model.BaseResponse{}

	url := ctx.Params("url")

	data, err := c.usecase.GetURL(url)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = data

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *url) FindAllUrl(ctx *fiber.Ctx) error {
	response := model.BaseResponse{}

	data, err := c.usecase.GetAllURL()
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Data = data

	return ctx.Status(fiber.StatusOK).JSON(response)
}