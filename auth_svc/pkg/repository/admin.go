package repository

import (
	"auth_svc/pkg/domain"
	interfaces "auth_svc/pkg/repository/interface"
	"context"
	"errors"

	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

// constructor implements admin interface return admin database struct

func NewadminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

// for findin admin
func (ad *adminDatabase) FindAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Where("username=? OR name = ? OR phone=? OR email=?", admin.Username, admin.Name, admin.Phone, admin.Email).First(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("admin not found")
	}
	return admin, nil
}

// for adding admin to database
func (ad *adminDatabase) AddAdmin(c context.Context, admin domain.AdminDetails) (domain.AdminDetails, error) {
	err := ad.DB.Create(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("error while adding admin details to database")
	}

	return admin, nil
}

func (ad *adminDatabase) FindByUsername(c context.Context, Username string) (domain.AdminDetails, error) {
	var admin domain.AdminDetails

	err := ad.DB.Raw("select *from admin_details where username=?", Username).Scan(&admin).Error
	if err != nil {
		return domain.AdminDetails{}, errors.New("failed find user details")
	}
	return admin, nil
}
