package interfaces

import (
	"auth_svc/pkg/domain"
	"context"
)

type AdminRepository interface {
	FindAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error)
	FindByUsername(c context.Context, Username string) (domain.AdminDetails, error)
}
