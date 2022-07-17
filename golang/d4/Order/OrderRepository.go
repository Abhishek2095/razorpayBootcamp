package Order

import (
	"ecommerce/Config"
	"errors"
)

type OrderRepository interface {
	GetOrderById(id string, order *Order) error
	GetAllOrders(orders *[]Order) error
	Save(order *Order) error
	GetOrderCount() int
	GetAllOrdersOfCustomer(id string, orders *[]Order) error
}

var (
	ErrOrderNotFound = errors.New("the order was not found in the repository")

	ErrFailedToAddOrder = errors.New("failed to add the order to the repository")

	ErrFailedToGetAllOrders = errors.New("failed to retrieve all orders from the repository")
)

type Repo struct{}

func NewRepo() OrderRepository {
	return &Repo{}
}

func (repo *Repo) GetAllOrdersOfCustomer(id string, orders *[]Order) error {
	if err := Config.DB.Where("customer_id = ?", id).Find(orders).Error; err != nil {
		return errors.New("no records found")
	}
	return nil
}

func (repo *Repo) GetOrderById(id string, order *Order) error {
	if err := Config.DB.Where("id = ?", id).Find(order).Error; err != nil {
		return ErrOrderNotFound
	}
	return nil
}

func (repo *Repo) GetAllOrders(orders *[]Order) error {
	if err := Config.DB.Find(orders).Error; err != nil {
		return ErrFailedToGetAllOrders
	}
	return nil
}

func (repo *Repo) Save(order *Order) error {
	if err := Config.DB.Save(order).Error; err != nil {
		return ErrFailedToAddOrder
	}
	return nil
}

func (repo *Repo) GetOrderCount() int {
	OrderCount := 0
	Config.DB.Model(&Order{}).Count(&OrderCount)
	return OrderCount
}
