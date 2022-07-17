package Services

import (
	"ecommerce/Customer"
	"ecommerce/Order"
	"ecommerce/Product"
	"ecommerce/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService_CreateOrder(t *testing.T) {
	asserter := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderRepo := mock.NewMockOrderRepository(ctrl)
	NewOrderService(orderRepo)

	productRepo := mock.NewMockProductRepository(ctrl)
	NewProductService(productRepo)

	customerRepo := mock.NewMockCustomerRepository(ctrl)
	NewCustomerService(customerRepo)

	customer := Customer.Customer{
		ID:      "CUST1",
		Name:    "test case",
		Address: "test address",
	}

	product := Product.Product{
		ID:       "PROD1",
		Name:     "bottle",
		Price:    100,
		Quantity: 4,
	}
	tc := Order.Order{
		ID:         "O1",
		CustomerId: "CUST1",
		ProductId:  "PROD1",
		Quantity:   3,
		Amount:     52,
	}

	expected := Order.Order{
		ID:         "ORD1",
		CustomerId: "CUST1",
		ProductId:  "PROD1",
		Quantity:   3,
		Amount:     300,
	}

	productRepo.EXPECT().GetProductById(gomock.Any(), gomock.Any()).SetArg(1, product).Return(nil).Times(2)
	customerRepo.EXPECT().GetCustomerById(gomock.Any(), gomock.Any()).SetArg(1, customer).Return(nil).Times(1)
	productRepo.EXPECT().Save(gomock.Any()).Return(nil).Times(1)
	orderRepo.EXPECT().GetOrderCount().Return(0).Times(1)
	orderRepo.EXPECT().Save(gomock.Any()).Return(nil).Times(1)

	err := OrderServices.CreateOrder(&tc)
	asserter.Nil(err)
	asserter.Equal(expected.ID, tc.ID)
	asserter.Equal(expected.ProductId, tc.ProductId)
	asserter.Equal(expected.CustomerId, tc.CustomerId)
	asserter.Equal(expected.Quantity, tc.Quantity)
	asserter.Equal(expected.Amount, tc.Amount)
}
