package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/pakawatkung/redis/repositories"
)

type productServiceRedis struct {
	proSrv repositories.ProductRepository
	redisClient *redis.Client
}

func NewProductServiceRedis(proSrv repositories.ProductRepository, redisClient *redis.Client) ProductService {
	return productServiceRedis{proSrv, redisClient}
}

func (s productServiceRedis) ServiceGetProduct() (produceRes []productResponse, err error)  {

	key := "service::getproduct"

	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		data := []productReceive{}
		if json.Unmarshal([]byte(productJson), &data) == nil {
			for _, values := range data {
				cleanData := productResponse{
					Name: values.Name,
					Qty: values.Qty,
				}
				produceRes = append(produceRes, cleanData)
			}
			return produceRes,nil
		}
	}

	data, err := s.proSrv.GetProduct()
	if err != nil {
		// should custom err
		return nil, err
	}
	
	for _, values := range data {
		cleanData := productResponse{
			Name: values.Name,
			Qty: values.Qty,
		}
		produceRes = append(produceRes, cleanData)
	}

	if data, err := json.Marshal(produceRes); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second * 10)
	}

	return produceRes, nil
}