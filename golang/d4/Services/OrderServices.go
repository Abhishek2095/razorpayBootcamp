package Services

import (
	"ecommerce/Customer"
	"ecommerce/DTO"
	"ecommerce/Order"
	"ecommerce/Product"
	mutex "ecommerce/Services/Mutex"
	"errors"
	"strconv"
)

var (
	OrderServices OrderServiceInterface

	ErrMissingCustomerId    = errors.New("missing customer id")
	ErrMissingProductId     = errors.New("missing product id")
	ErrInvalidQuantity      = errors.New("invalid order quantity")
	ErrQuantityNotAvailable = errors.New("quantity not available")
)

type OrderServiceInterface interface {
	CreateOrder(Order *Order.Order) error
	GetOrderById(id string, product *Order.Order) error
	CreateMultipleOrders(Orders *DTO.MultipleOrder) error
	GetCustomerOrderHistory(id string, orders *[]Order.Order) error
	GetOrderHistory(orders *[]Order.Order) error
}

type OrderService struct {
	repo Order.OrderRepository
}

func NewOrderService(repository Order.OrderRepository) {
	OrderServices = &OrderService{repo: repository}
}

func (orderService *OrderService) GetOrderHistory(orders *[]Order.Order) error {
	if err := orderService.repo.GetAllOrders(orders); err != nil {
		return err
	}
	return nil
}

func (orderService *OrderService) GetCustomerOrderHistory(id string, orders *[]Order.Order) error {
	if err := orderService.repo.GetAllOrdersOfCustomer(id, orders); err != nil {
		return err
	}
	return nil
}

func (orderService *OrderService) CreateOrder(Order *Order.Order) error {
	if Order.CustomerId == "" {
		return ErrMissingCustomerId
	}
	if Order.ProductId == "" {
		return ErrMissingProductId
	}
	if Order.Quantity < 1 {
		return ErrInvalidQuantity
	}

	product := Product.Product{}
	if err := ProductServices.GetProductById(Order.ProductId, &product); err != nil {
		return err
	}

	customer := Customer.Customer{}
	if err := CustomerServices.GetCustomerById(Order.CustomerId, &customer); err != nil {
		return err
	}

	mutex.Mutex.Lock(Order.ProductId)
	defer mutex.Mutex.UnLock(Order.ProductId)

	if product.Quantity < Order.Quantity {
		return ErrQuantityNotAvailable
	} else {
		product.Quantity -= Order.Quantity
		ProductServices.UpdateProduct(&product)
	}

	OrderCount := orderService.repo.GetOrderCount() + 1
	Order.ID = "ORD" + strconv.Itoa(OrderCount)
	Order.Amount = product.Price * Order.Quantity

	if err := orderService.repo.Save(Order); err != nil {
		return err
	}
	return nil
}

func (orderService *OrderService) GetOrderById(id string, Order *Order.Order) error {
	if err := orderService.repo.GetOrderById(id, Order); err != nil {
		return err
	}
	return nil
}

func (orderService *OrderService) CreateMultipleOrders(Orders *DTO.MultipleOrder) error {
	for i := range Orders.ProductID {
		order := Order.Order{
			CustomerId: Orders.CustomerID,
			ProductId:  Orders.ProductID[i],
			Quantity:   Orders.Quantity[i],
		}
		orderService.CreateOrder(&order)
	}
	return nil
}
