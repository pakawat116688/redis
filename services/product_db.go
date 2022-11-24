package services

import "github.com/pakawatkung/redis/repositories"

type productService struct {
	proSrv repositories.ProductRepository
}

func NewProductService(proSrv repositories.ProductRepository) ProductService {
	return productService{proSrv: proSrv}
}

func (s productService) ServiceGetProduct() ([]productResponse, error)  {
	
	return nil, nil
}