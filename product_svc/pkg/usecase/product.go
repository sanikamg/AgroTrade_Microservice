package usecase

import (
	"context"
	"errors"
	"product_svc/pkg/domain"
	interfaces "product_svc/pkg/repository/interface"
	ser "product_svc/pkg/usecase/interface"
	"product_svc/pkg/utils"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(repo interfaces.ProductRepository) ser.ProductUsecase {
	return &ProductUsecase{
		productRepo: repo,
	}
}

// product
func (pu *ProductUsecase) AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {

	produ, err := pu.productRepo.FindProduct(c, product)
	product.Product_Id = produ.Product_Id
	if err == nil {
		// prod, err := pu.productRepo.AddQuantity(c, product)
		// if err != nil {
		// 	return domain.ProductDetails{}, err
		// }

		// return prod, nil
		return produ, errors.New("product already exist please update product")
	}

	pro, err := pu.productRepo.AddProduct(c, product)
	if err != nil {
		return domain.ProductDetails{}, err
	}
	return pro, nil
}

func (pu *ProductUsecase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error) {
	product, metadata, err := pu.productRepo.FindAllProducts(c, pagination)
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	return product, metadata, nil
}
