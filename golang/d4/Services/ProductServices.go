package Services

import (
	"ecommerce/Product"
	"errors"
	"strconv"
)

var (
	ProductServices ProductServiceInterface

	ErrMissingProductName     = errors.New("missing product name")
	ErrInvalidProductQuantity = errors.New("invalid product quantity")
	ErrInvalidProductPrice    = errors.New("invalid product price")
)

type ProductServiceInterface interface {
	CreateProduct(product *Product.Product) error
	GetAllProducts(products *[]Product.Product) error
	GetProductById(id string, product *Product.Product) error
	UpdateProduct(product *Product.Product) error
}

type ProductService struct {
	repo Product.ProductRepository
}

func NewProductService(repository Product.ProductRepository) {
	ProductServices = &ProductService{repo: repository}
}

func (productService *ProductService) CreateProduct(product *Product.Product) error {
	if product.Name == "" {
		return ErrMissingProductName
	}
	if product.Price < 1 {
		return ErrInvalidProductPrice
	}
	if product.Quantity < 1 {
		return ErrInvalidProductQuantity
	}

	productCount := productService.repo.GetProductCount()
	productCount++
	product.ID = "PROD" + strconv.Itoa(productCount)

	if err := productService.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (productService *ProductService) UpdateProduct(product *Product.Product) error {
	if product.Name == "" {
		return ErrMissingProductName
	}
	if product.Price < 1 {
		return ErrInvalidProductPrice
	}

	productExists := Product.Product{}
	if err := productService.GetProductById(product.ID, &productExists); err != nil {
		return err
	}

	if err := productService.repo.Save(product); err != nil {
		return err
	}
	return nil
}

func (productService *ProductService) GetAllProducts(products *[]Product.Product) error {
	if err := productService.repo.GetAllProducts(products); err != nil {
		return err
	}
	return nil
}

func (productService *ProductService) GetProductById(id string, product *Product.Product) error {
	if err := productService.repo.GetProductById(id, product); err != nil {
		return err
	}
	return nil
}
