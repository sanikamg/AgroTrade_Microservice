package repository

import (
	"auth_svc/pkg/domain"
	interfaces "auth_svc/pkg/repository/interface"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// constructor implement UserRepository interface return userDatabase struct

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>user signup>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.....
func (ud *userDatabase) Addusers(ctx context.Context, user domain.Users) (domain.Users, error) {

	err := ud.DB.Create(&user).Error

	if err != nil {
		return domain.Users{}, fmt.Errorf("error adding users: %w", err)
	}
	return user, nil
}

// to check user already exist or not
func (ud *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := ud.DB.Where("username = ?  OR email = ?", user.Username, user.Email).First(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("user not found")
	}

	return user, nil
}

// find user by phonenumber
func (ud *userDatabase) FindUserByPhn(c context.Context, phn domain.Users) error {
	err := ud.DB.Where("phone=?", phn.Phone).First(&phn).Error
	if err != nil {
		return errors.New("failed to find user")
	}
	return nil
}

// finding username is empty
func (ud *userDatabase) IsEmtyUsername(c context.Context, username domain.Users) bool {
	ud.DB.Where("phone = ?", username.Phone).First(&username)
	return username.Username == ""
}

//update user status

func (ud *userDatabase) UpdateStatus(c context.Context, user domain.Users) error {
	query := `update users set verification=? where phone=?`
	err := ud.DB.Raw(query, true, user.Phone).Scan(&user).Error
	if err != nil {
		return errors.New("failed to update update status")
	}
	return nil
}

//find status to update details

func (ud *userDatabase) FindStatus(c context.Context, phn string) (domain.Users, error) {
	var usr domain.Users
	query := `select *from users where phone=?`
	err := ud.DB.Raw(query, phn).Scan(&usr).Error
	if err != nil {
		return domain.Users{}, errors.New("failed to find status")
	}
	return usr, nil

}

// complete user profile
func (ud *userDatabase) UpdateUserDetails(c context.Context, user domain.Users) (domain.Users, error) {

	query := `update users set username=?,name=?,email=?,password=? where phone=?`
	err := ud.DB.Raw(query, user.Username, user.Name, user.Email, user.Password, user.Phone).Scan(&user).Error
	if err != nil {
		return domain.Users{}, errors.New("failed to complete user registration")
	}
	return user, nil
}
