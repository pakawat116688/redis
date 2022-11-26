package handlers

import "github.com/gofiber/fiber/v2"

type productHandlers interface {
	GetProductAPI(c *fiber.Ctx) error
}