package services

type productResponse struct {
	Name string `json:"name"`
	Qty int `json:"qty"`
}

type ProductService interface {
	ServiceGetProduct() ([]productResponse, error)
}