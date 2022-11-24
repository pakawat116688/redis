package repositories

import (
	"gorm.io/gorm"
)

type procductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return procductDB{db: db}
}

func (r procductDB) GetProduct() (product []product, err error)  {
	err = r.db.Order("qty desc").Find(&product).Error
	return product, err
}