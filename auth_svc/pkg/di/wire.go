//go:build wireinject
// +build wireinject

package di

import (
	"auth_svc/pkg/config"

	"auth_svc/pkg/db"
	"auth_svc/pkg/repository"
	svc "auth_svc/pkg/service"
	"auth_svc/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) svc.AuthService {
	wire.Build(
		db.ConnectDatabase,
		//repository
		repository.NewadminRepository, repository.NewUserRepository,

		//usecase
		usecase.NewadminUsecase, usecase.NewUserUsecase,

		//service
		svc.NewAuthService,
	)
	return svc.AuthService{}
}
