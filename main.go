package main

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/redis/handlers"
	"github.com/pakawatkung/redis/repositories"
	"github.com/pakawatkung/redis/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main()  {

	dsn := "root:passw0rd@tcp(mariadb-service)/product?parseTime=true"
	director := mysql.Open(dsn)
	db, err := gorm.Open(director)
	if err != nil {
		panic(err)
	}

	myRedis := redis.NewClient(&redis.Options{
		Addr: "redis-service:6379",
	})

	redisRepo := repositories.NewProductRedis(db,myRedis)
	redisSrv := services.NewProductServiceRedis(redisRepo, myRedis)
	redisHanle := handlers.NewProductHandleRedis(redisSrv, myRedis)

	productRepo := repositories.NewProductDB(db)
	productSrv := services.NewProductService(productRepo)
	productHandle := handlers.NewProductHandler(productSrv)


	app := fiber.New()

	app.Get("/products", productHandle.GetProductAPI)
	app.Get("/products/redis", redisHanle.GetProductAPI)

	app.Listen(":8000")
}