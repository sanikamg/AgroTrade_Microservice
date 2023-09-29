//go:build wireinject
// +build wireinject

package di

import (
	"product_svc/pkg/config"
	"product_svc/pkg/db"
	"product_svc/pkg/repository"
	svc "product_svc/pkg/service"
	"product_svc/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (svc.ProductService, error) {
	wire.Build(
		db.ConnectDatabase,
		//repository
		repository.NewProductRepository,

		//usecase
		usecase.NewProductUsecase,

		//service
		svc.NewProductService,
	)
	return svc.ProductService{}, nil
}
