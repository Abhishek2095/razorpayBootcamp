package Customer

import (
	"ecommerce/Config"
	"errors"
)

type CustomerRepository interface {
	GetCustomerById(id string, customer *Customer) error
	Save(customer *Customer) error
	GetCustomerCount() int
}

var (
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")

	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
)

type Repo struct{}

func NewRepo() CustomerRepository {
	return &Repo{}
}

func (repo *Repo) GetCustomerById(id string, customer *Customer) error {
	if err := Config.DB.Where("id = ?", id).Find(customer).Error; err != nil {
		return ErrFailedToAddCustomer
	}
	return nil
}

func (repo *Repo) Save(customer *Customer) error {
	if err := Config.DB.Save(customer).Error; err != nil {
		return ErrFailedToAddCustomer
	}
	return nil
}

func (repo *Repo) GetCustomerCount() int {
	var CustomerCount = 0
	Config.DB.Model(&Customer{}).Count(&CustomerCount)
	return CustomerCount
}
