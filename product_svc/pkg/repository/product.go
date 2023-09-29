package repository

import (
	"context"
	"errors"
	"fmt"
	"product_svc/pkg/domain"
	interfaces "product_svc/pkg/repository/interface"
	"product_svc/pkg/utils"

	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

// constructor implements admin interface return admin database struct
// product
func (pd *productDatabase) AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {
	err := pd.DB.Create(&product).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to add product")
	}
	return product, nil
}

// find
func (pd *productDatabase) FindProductById(c context.Context, productid uint) error {
	var product domain.ProductDetails
	err := pd.DB.Where("product_id=?", productid).First(&product).Error
	if err != nil {
		return errors.New("failed to find product")
	}
	return nil
}

func (pd *productDatabase) FindProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {
	err := pd.DB.Where("product_id=? OR product_name=?", product.Product_Id, product.ProductName).First(&product).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to find product")
	}

	return product, nil
}

func (pd *productDatabase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.ProductDetails{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}

	query := `SELECT p.product_id, p.product_name, p.product_price, p.product_quantity, p.category
	FROM product_details AS p LIMIT $1 OFFSET $2;
	`
	rows, err := db.Raw(query, pagination.Limit(), pagination.Offset()).Rows()
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	defer rows.Close()
	fmt.Println(products)
	for rows.Next() {
		var product domain.ProductResponse

		err := rows.Scan(
			&product.Product_Id,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductQuantity,
			&product.Category_name,
		)
		if err != nil {
			return []domain.ProductResponse{}, utils.Metadata{}, err
		}
		fmt.Println(product)
		products = append(products, product)
	}

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return products, metadata, nil
}
