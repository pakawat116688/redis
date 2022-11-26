package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/redis/services"
)

type productHandler struct {
	productHandle services.ProductService
}

func NewProductHandler(productHandle services.ProductService) productHandlers  {
	return productHandler{productHandle}
}

func (h productHandler) GetProductAPI(c *fiber.Ctx) error  {

	products, err := h.productHandle.ServiceGetProduct()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	response := fiber.Map{
		"status": "OK",
		"path": "/product_api",
		"product": products,
	}
	return c.Status(fiber.StatusOK).JSON(response)

}