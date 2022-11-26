package services

type productReceive struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type productResponse struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type ProductService interface {
	ServiceGetProduct() ([]productResponse, error)
}
