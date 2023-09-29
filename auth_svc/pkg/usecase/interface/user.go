package interfaces

import (
	"auth_svc/pkg/domain"
	"context"
)

type UserUsecase interface {
	//user signup
	Register(ctx context.Context, user domain.Users) (domain.Users, error)
	//Adduser(ctx context.Context, user domain.Users) (domain.Users, error)
	UpdateStatus(c context.Context, user domain.Users) error
	SendOtpPhn(c context.Context, phn domain.Users) error
	VerifyOtp(c context.Context, phn string, otp string) error
	//user login
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
}
