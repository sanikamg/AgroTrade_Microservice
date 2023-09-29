package usecase

import (
	"auth_svc/pkg/domain"
	interfaces "auth_svc/pkg/repository/interface"
	ser "auth_svc/pkg/usecase/interface"
	"auth_svc/pkg/verification"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(repo interfaces.UserRepository) ser.UserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
} 
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>login/sign up>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// func generateOTPCode() string {
// 	// Generate a random 6-digit OTP code
// 	return fmt.Sprintf("%06d", rand.Intn(1000000))
// }

func (uu *UserUsecase) SendOtpPhn(c context.Context, phn domain.Users) error {
	err := uu.userRepo.FindUserByPhn(c, phn)
	if err == nil {
		if uu.userRepo.IsEmtyUsername(c, phn) {
			return errors.New("user verification already completed please complete registration")
		}
		return errors.New("user already exists please login")
	}
	// Generate OTP code

	if _, err1 := verification.SendOtp("+91" + phn.Phone); err1 != nil {

		return errors.Join(errors.New("can't send otp"), err1)
	}

	return nil
}

// func (c *UserUsecase) Adduser(ctx context.Context, user domain.Users) (domain.Users, error) {

// 	c.userRepo.Addusers(ctx, user)

// 	return user, nil

// }

//verify otp

func (uu *UserUsecase) VerifyOtp(c context.Context, phn string, otp string) error {
	var usr domain.Users
	err := verification.VerifyOtp("+91"+phn, otp)
	if err != nil {
		return errors.New("failed to verify otp")
	}
	usr.Phone = phn
	_, er := uu.userRepo.Addusers(c, usr)
	if er != nil {
		return errors.New("can't add user ")
	}
	return nil
}

func (uu *UserUsecase) UpdateStatus(c context.Context, user domain.Users) error {

	err := uu.userRepo.UpdateStatus(c, user)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {
	//to hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return domain.Users{}, errors.New("bcrypt failed:" + err.Error())
	}
	user.Password = string(hash)
	usr, err := uu.userRepo.FindStatus(ctx, user.Phone)
	if err != nil {
		return domain.Users{}, err
	}
	user.Verification = usr.Verification
	user.User_Id = usr.User_Id

	if user.Verification {
		usr, err := uu.userRepo.UpdateUserDetails(ctx, user)
		if err != nil {
			return domain.Users{}, err
		}

		return usr, nil
	}
	return domain.Users{}, errors.New("enter correct details")
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.login>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (c *UserUsecase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	dbUser, dbErr := c.userRepo.FindUser(ctx, user)

	//check whether the user exists or valid information
	if dbErr != nil {
		return domain.Users{}, dbErr
	} else if dbUser.User_Id == 0 {
		return domain.Users{}, errors.New("user does not exists with this , please register")
	}

	//checking block status

	if dbUser.BlockStatus {
		return domain.Users{}, errors.New("blocked user trying to login")
	}

	// check password matching

	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return domain.Users{}, errors.New("password is not correct")
	}

	return dbUser, nil
}
