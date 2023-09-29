package interfaces

import (
	"auth_svc/pkg/domain"
	"context"
)

type UserRepository interface {
	Addusers(ctx context.Context, user domain.Users) (domain.Users, error)
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
	IsEmtyUsername(c context.Context, username domain.Users) bool
	UpdateStatus(c context.Context, user domain.Users) error
	FindStatus(c context.Context, phn string) (domain.Users, error)
	FindUserByPhn(c context.Context, phn domain.Users) error
	UpdateUserDetails(c context.Context, user domain.Users) (domain.Users, error)
}
