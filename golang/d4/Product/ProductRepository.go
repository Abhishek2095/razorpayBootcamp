package Product

import (
	"ecommerce/Config"
	"errors"
)

type ProductRepository interface {
	GetProductById(id string, product *Product) error
	GetAllProducts(products *[]Product) error
	Save(product *Product) error
	GetProductCount() int
}

var (
	ErrProductNotFound = errors.New("the product was not found in the repository")

	ErrFailedToAddProduct = errors.New("failed to add the product to the repository")

	ErrFailedToGetAllProducts = errors.New("failed to get products from the repository")
)

type Repo struct{}

func NewRepo() ProductRepository {
	return &Repo{}
}

func (repo *Repo) GetProductById(id string, product *Product) error {
	if err := Config.DB.Where("id = ?", id).Find(product).Error; err != nil {
		return ErrProductNotFound
	}
	return nil
}

func (repo *Repo) Save(product *Product) error {
	if err := Config.DB.Save(product).Error; err != nil {
		return ErrFailedToAddProduct
	}
	return nil
}

func (repo *Repo) GetAllProducts(products *[]Product) error {
	if err := Config.DB.Find(products).Error; err != nil {
		return ErrFailedToGetAllProducts
	}
	return nil
}

func (repo *Repo) GetProductCount() int {
	var ProductCount = 0
	Config.DB.Model(&Product{}).Count(&ProductCount)
	return ProductCount
}
