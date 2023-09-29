package interfaces

import (
	"auth_svc/pkg/domain"
	"context"
)

type AdminUsecase interface {
	AdminSignup(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
	AdminLogin(ctx context.Context, admin domain.AdminDetails) error
}
