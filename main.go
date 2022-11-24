package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/redis/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {

	dsn := "root:passw0rd@tcp(localhost)/product?parseTime=true"
	director := mysql.Open(dsn)
	db, err := gorm.Open(director)
	if err != nil {
		panic(err)
	}

	productRepo := repositories.NewProductDB(db)
	_ = productRepo

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":8000")
}