package services

import "github.com/pakawatkung/redis/repositories"

type productService struct {
	proSrv repositories.ProductRepository
}

func NewProductService(proSrv repositories.ProductRepository) ProductService {
	return productService{proSrv: proSrv}
}

func (s productService) ServiceGetProduct() (productRes []productResponse, err error)  {

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
		productRes = append(productRes, cleanData)
	}

	return productRes, nil
}