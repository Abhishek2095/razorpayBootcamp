package Services

import (
	"ecommerce/Product"
	"ecommerce/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService_CreateProduct(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockProductRepository(ctrl)
	NewProductService(repo)

	tc := Product.Product{
		Name:     "bottle",
		Price:    100,
		Quantity: 5,
	}

	expected := Product.Product{
		ID:       "PROD1",
		Name:     "bottle",
		Price:    100,
		Quantity: 5,
	}

	repo.EXPECT().Save(&tc).Return(nil).Times(1)
	repo.EXPECT().GetProductCount().Return(0).Times(1)

	err := ProductServices.CreateProduct(&tc)
	assertion.Nil(err)
	assertion.Equal(expected.ID, tc.ID)
	assertion.Equal(expected.Name, tc.Name)
	assertion.Equal(expected.Price, tc.Price)
	assertion.Equal(expected.Quantity, tc.Quantity)
}

func TestProductService_UpdateProduct(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockProductRepository(ctrl)
	NewProductService(repo)

	tc := Product.Product{
		Name:     "bottle",
		Price:    100,
		Quantity: 5,
	}

	expected := Product.Product{
		ID:       "PROD1",
		Name:     "bottle",
		Price:    100,
		Quantity: 4,
	}

	repo.EXPECT().Save(&tc).Return(nil).Times(1)
	repo.EXPECT().GetProductCount().Return(0).Times(1)

	err := ProductServices.CreateProduct(&tc)
	assertion.Nil(err)

	repo.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	repo.EXPECT().Save(&tc).Return(nil).Times(1)

	tc.Quantity -= 1

	err = ProductServices.UpdateProduct(&tc)
	assertion.Nil(err)
	assertion.Equal(expected.ID, tc.ID)
	assertion.Equal(expected.Name, tc.Name)
	assertion.Equal(expected.Price, tc.Price)
	assertion.Equal(expected.Quantity, tc.Quantity)
}

func TestProductService_GetProductById(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockProductRepository(ctrl)
	NewProductService(repo)

	tc := Product.Product{
		Name:     "",
		Price:    2,
		Quantity: 1,
	}

	expected := Product.Product{
		Name:     "bottle",
		Price:    100,
		Quantity: 5,
	}

	repo.EXPECT().Save(&expected).Return(nil).Times(1)
	repo.EXPECT().GetProductCount().Return(0).Times(1)

	err := ProductServices.CreateProduct(&expected)
	assertion.Nil(err)

	repo.EXPECT().GetProductById(gomock.Any(), gomock.Any()).SetArg(1, Product.Product{
		ID:       "PROD1",
		Name:     "bottle",
		Price:    100,
		Quantity: 5,
	}).Times(1)

	err = ProductServices.GetProductById(tc.ID, &tc)
	assertion.Nil(err)
	assertion.Equal(expected.ID, tc.ID)
	assertion.Equal(expected.Name, tc.Name)
	assertion.Equal(expected.Price, tc.Price)
	assertion.Equal(expected.Quantity, tc.Quantity)
}
