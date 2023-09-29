package service

import (
	"auth_svc/pkg/auth/pb"
	"auth_svc/pkg/domain"
	service "auth_svc/pkg/usecase/interface"
	"auth_svc/pkg/verification"
	"context"
)

type AuthService struct {
	adminUsecase service.AdminUsecase
	userUsecase  service.UserUsecase
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(a service.AdminUsecase, u service.UserUsecase) AuthService {
	return AuthService{
		adminUsecase: a,
		userUsecase:  u,
	}
}

var admin domain.AdminDetails

func (a *AuthService) AdminSignup(ctx context.Context, req *pb.AdminDetailsRequest) (*pb.Response, error) {

	if _, err := verification.SendOtp("+91" + req.Phone); err != nil {
		return &pb.Response{
			Message:    "sending otp failed",
			Statuscode: 400,
			Errors:     "sending otp failed",
		}, nil
	}

	admin.Email = req.Email
	admin.Name = req.Name
	admin.Username = req.Username
	admin.Password = req.Password
	admin.Phone = req.Phone

	return &pb.Response{
		Message:    "otp send succesfully",
		Statuscode: 200,
	}, nil
}

func (a *AuthService) VerifyOtp(ctx context.Context, req *pb.OTPRequest) (*pb.Response, error) {

	err1 := verification.VerifyOtp("+91"+req.Phone, req.OTP)

	if err1 != nil {
		return &pb.Response{
			Message:    "otp verification failed",
			Statuscode: 400,
			Errors:     "otp can'be verified ",
		}, nil
	}

	_, err := a.adminUsecase.AdminSignup(ctx, admin)

	if err != nil {
		return &pb.Response{
			Message:    "otp verification failed",
			Statuscode: 400,
			Errors:     "otp verification failed",
		}, nil
	}

	return &pb.Response{
		Message:    "otp send succesfully",
		Statuscode: 200,
	}, nil
}

func (a *AuthService) AdminLogin(ctx context.Context, req *pb.LoginDetailsRequest) (*pb.AdminResponse, error) {
	// check any field is empty
	if req.Username == " " && req.Password == " " {
		return &pb.AdminResponse{
			Message:    "enter your correct username and password",
			Statuscode: 400,
			Errors:     "username and password didn't match ",
		}, nil

	}
	admin.Password = req.Password
	admin.Username = req.Username

	// admin, err := a.adminUsecase.FindByUsername(ctx, req.Username)
	// if err != nil {
	// 	return &pb.AdminResponse{
	// 		Message:    "Enter valid username",
	// 		Statuscode: 400,
	// 		Errors:     "login failed ",
	// 	}, nil
	// }

	// check whether the user exists and login usisng usecse function
	if err := a.adminUsecase.AdminLogin(ctx, admin); err != nil {
		return &pb.AdminResponse{
			Message:    "Enter valid username",
			Statuscode: 400,
			Errors:     "login failed ",
		}, nil
	}

	return &pb.AdminResponse{
		Message:    "succesfully logged in",
		Statuscode: 200,
		Errors:     " ",
		ID:         uint32(admin.ID),
	}, nil

}
