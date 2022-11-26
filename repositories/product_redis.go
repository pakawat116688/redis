package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type productRedis struct {
	db *gorm.DB
	redisClient *redis.Client
}

func NewProductRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository  {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRedis{db, redisClient}
}

func (r productRedis) GetProduct() (products []product, err error)  {

	key := "repository::getproduct"
	productjson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productjson), &products)
		if err == nil {
			return products, nil
		}
	}

	err = r.db.Order("qty desc").Find(&products).Error
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), key, string(data), time.Second * 10).Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}