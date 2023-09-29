package interfaces

import (
	"context"
	"product_svc/pkg/domain"
	"product_svc/pkg/utils"
)

type ProductRepository interface {
	AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindProductById(c context.Context, productid uint) error
	FindProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error)
	FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error)
}
