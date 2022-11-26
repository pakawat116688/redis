package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/redis/services"
)

type productHandleRedis struct {
	productHandle services.ProductService
	redisClient *redis.Client
}

func NewProductHandleRedis(productHandle services.ProductService, redisClient *redis.Client) productHandlers {
	return productHandleRedis{productHandle, redisClient}
}

func (h productHandleRedis) GetProductAPI(c *fiber.Ctx) error {

	key := "handle::getproduct"
	// Get Redis
	if product, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("from redis")
		c.Set("Content-Type", "application/json")
		return c.Status(fiber.StatusOK).SendString(product)
	}

	// Service
	products, err := h.productHandle.ServiceGetProduct()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	fmt.Println("from database")
	response := fiber.Map{
		"status": "OK",
		"path": "/product_redis_api",
		"product": products,
	}

	// Redis SET
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second * 10)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}