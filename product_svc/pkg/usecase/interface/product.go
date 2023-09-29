package interfaces

import (
	"context"
	"product_svc/pkg/domain"
	"product_svc/pkg/utils"
)

type ProductUsecase interface {
	//product
	AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error)
}
