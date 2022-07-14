package Services

import (
	"ecommerce/Customer"
	"errors"
	"strconv"
)

var (
	CustomerServices CustomerServiceInterface

	ErrMissingAddress = errors.New("missing address")
	ErrMissingName    = errors.New("missing name")
)

type CustomerServiceInterface interface {
	CreateCustomer(customer *Customer.Customer) error
	GetCustomerById(id string, product *Customer.Customer) error
}

type CustomerService struct {
	repo Customer.CustomerRepository
}

func NewCustomerService(repository Customer.CustomerRepository) {
	CustomerServices = &CustomerService{repo: repository}
}

func (customerService *CustomerService) CreateCustomer(customer *Customer.Customer) error {
	if customer.Name == "" {
		return ErrMissingName
	}
	if customer.Address == "" {
		return ErrMissingAddress
	}

	customerCount := customerService.repo.GetCustomerCount() + 1
	customer.ID = "CUST" + strconv.Itoa(customerCount)

	if err := customerService.repo.Save(customer); err != nil {
		return err
	}
	return nil
}

func (customerService *CustomerService) GetCustomerById(id string, customer *Customer.Customer) error {
	if err := customerService.repo.GetCustomerById(id, customer); err != nil {
		return err
	}
	return nil
}
