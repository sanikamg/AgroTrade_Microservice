package service

import (
	"auth_svc/pkg/auth/pb"
	"auth_svc/pkg/domain"
	"auth_svc/pkg/verification"
	"context"
	"fmt"
)

var user domain.Users

func (u *AuthService) UserSignup(c context.Context, req *pb.PhoneRequest) (*pb.Response, error) {
	fmt.Println(req)
	var user domain.Users
	if _, err := verification.SendOtp("+91" + req.Phone); err != nil {
		return &pb.Response{
			Message:    "sending otp failed",
			Statuscode: 400,
			Errors:     "sending otp failed- " + err.Error(),
		}, nil
	}

	user.Phone = req.Phone

	return &pb.Response{
		Message:    "otp send successfully",
		Statuscode: 200,
	}, nil
}

// verify otp
func (u *AuthService) UserVerifyOtp(ctx context.Context, req *pb.OTPRequest) (*pb.Response, error) {

	// verifying otp

	err1 := u.userUsecase.VerifyOtp(ctx, req.Phone, req.OTP)

	if err1 != nil {
		return &pb.Response{
			Message:    "otp verification failed",
			Statuscode: 400,
			Errors:     "otp can't verified",
		}, nil
	}

	user.Phone = req.Phone

	err2 := u.userUsecase.UpdateStatus(ctx, user)

	if err2 != nil {
		return &pb.Response{
			Message:    "status updation  failed",
			Statuscode: 400,
			Errors:     "status updation error",
		}, nil
	}

	return &pb.Response{
		Message:    "otp send successfully",
		Statuscode: 200,
	}, nil
}

func (u *AuthService) Register(ctx context.Context, req *pb.AdminDetailsRequest) (*pb.Response, error) {

	user.Email = req.Email
	user.Name = req.Name
	user.Username = req.Username
	user.Password = req.Password
	user.Phone = req.Phone

	_, err := u.userUsecase.Register(ctx, user)
	if err != nil {
		return &pb.Response{
			Message:    "sending otp failed",
			Statuscode: 400,
			Errors:     "sending otp failed",
		}, nil
	}

	return &pb.Response{
		Message:    "registration completed",
		Statuscode: 200,
	}, nil

}

func (u *AuthService) UserLogin(ctx context.Context, req *pb.LoginDetailsRequest) (*pb.AdminResponse, error) {
	// check any field is empty
	if req.Username == " " && req.Password == " " {
		return &pb.AdminResponse{
			Message:    "enter your correct username and password",
			Statuscode: 400,
			Errors:     "username and password didn't match ",
		}, nil

	}
	user.Password = req.Password
	user.Username = req.Username

	// check whether the user exists and login usisng usecse function
	if _, err := u.userUsecase.Login(ctx, user); err != nil {
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
		ID:         uint32(user.User_Id),
	}, nil

}
