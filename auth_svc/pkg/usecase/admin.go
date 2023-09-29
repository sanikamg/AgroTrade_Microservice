package usecase

import (
	"auth_svc/pkg/domain"
	interfaces "auth_svc/pkg/repository/interface"
	ser "auth_svc/pkg/usecase/interface"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct {
	adminRepo interfaces.AdminRepository
}

func NewadminUsecase(repo interfaces.AdminRepository) ser.AdminUsecase {
	return &AdminUsecase{
		adminRepo: repo,
	}
}

// AdminSignup implements interfaces.AdminUsecase
func (ad *AdminUsecase) AdminSignup(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {

	adn, err := ad.adminRepo.FindAdmin(c, admin)
	if err == nil {
		return adn, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while hashing")
	}
	admin.Password = string(hash)

	ad.adminRepo.AddAdmin(c, admin)

	return admin, nil

}

// find admin details by username
func (ad *AdminUsecase) FindByUsername(c context.Context, Username string) (domain.AdminDetails, error) {
	admin, err := ad.adminRepo.FindByUsername(c, Username)
	if err != nil {
		return domain.AdminDetails{}, err
	}
	return admin, nil
}

func (ad *AdminUsecase) AdminLogin(ctx context.Context, admin domain.AdminDetails) error {
	dbAdmin, dbErr := ad.adminRepo.FindAdmin(ctx, admin)
	fmt.Println(dbAdmin)
	fmt.Println(admin)

	//check whether the user exists or valid information
	if dbErr == nil {
		return dbErr
	} else if dbAdmin.ID == 0 {
		return errors.New("user does not exists with this , please register")
	}

	// check password matching

	if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
		return errors.New("password is not correct")
	}
	
	return nil
}
