package Services

import (
	"ecommerce/Customer"
	"ecommerce/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerService_CreateCustomer(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockCustomerRepository(ctrl)
	NewCustomerService(repo)

	tt := Customer.Customer{
		Name:    "test case",
		Address: "test address",
	}

	expected := Customer.Customer{
		ID:      "CUST1",
		Name:    "test case",
		Address: "test address",
	}

	repo.EXPECT().Save(&tt).Return(nil).Times(1)
	repo.EXPECT().GetCustomerCount().Return(0).Times(1)

	err := CustomerServices.CreateCustomer(&tt)
	assertion.Nil(err)
	assertion.Equal(expected.ID, tt.ID)
	assertion.Equal(expected.Name, tt.Name)
	assertion.Equal(expected.Address, tt.Address)
}

// func TestCustomerService_GetCustomerById(t *testing.T) {
// 	asserter := assert.New(t)

// 	mockController := gomock.NewController(t)
// 	defer mockController.Finish()

// 	repo := mock.NewMockCustomerRepository(mockController)
// 	NewCustomerService(repo)

// 	tc := Customer.Customer{
// 		ID: "CUST1",
// 	}

// 	expected := Customer.Customer{
// 		Name:    "test case",
// 		Address: "test address",
// 	}

// 	repo.EXPECT().GetCustomerCount().Return(0).Times(1)
// 	repo.EXPECT().Save(&tc).Return(nil).Times(1)
// 	repo.EXPECT().GetCustomerById(tc.ID, &tc).Return().Times(1)

// 	err := CustomerServices.CreateCustomer(&expected)
// 	asserter.Nil(err)

// 	err = CustomerServices.GetCustomerById(tc.ID, &tc)
// 	asserter.Nil(err)
// 	asserter.Equal(expected.ID, tc.ID)
// 	asserter.Equal(expected.Name, tc.Name)
// 	asserter.Equal(expected.Address, tc.Address)
// }
