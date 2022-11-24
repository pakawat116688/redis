package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type product struct {
	ID   int    `gorm:"column:pid"`
	Name string `gorm:"column:name"`
	Qty  int    `gorm:"column:qty"`
}

type ProductRepository interface {
	GetProduct() ([]product, error)
}

func mockData(db *gorm.DB) error  {

	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}
	
	randTemp := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randTemp)

	products := []product{}
	for i := 0; i < 30; i++ {
		p := product{
			ID: i,
			Name: fmt.Sprintf("Coca-Cola %v",i+1),
			Qty: random.Intn(100),
		}
		products = append(products, p)
	}
	
	return db.Create(&products).Error

}