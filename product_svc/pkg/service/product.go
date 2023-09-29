package service

import (
	"context"
	"fmt"
	"product_svc/pkg/domain"
	"product_svc/pkg/product/pb"
	service "product_svc/pkg/usecase/interface"
	"product_svc/pkg/utils"
)

type ProductService struct {
	productUsecase service.ProductUsecase
	pb.UnimplementedProductServiceServer
}

func NewProductService(p service.ProductUsecase) ProductService {
	return ProductService{

		productUsecase: p,
	}
}

func (p *ProductService) AddProduct(ctx context.Context, req *pb.ProductDetailsRequest) (*pb.Response, error) {
	var product domain.ProductDetails
	product.ProductName = req.ProductName
	product.ProductPrice = uint(req.ProductPrice)
	product.ProductQuantity = uint(req.ProductQuanity)
	product.Category = req.Category

	_, err := p.productUsecase.AddProduct(ctx, product)
	if err != nil {
		return &pb.Response{
			Message:    "product added failed",
			Statuscode: 400,
			Errors:     "can't be add product",
		}, nil
	}
	return &pb.Response{
		Message:    "product succesfully added",
		Statuscode: 200,
	}, nil

}

func (p *ProductService) GetProduct(ctx context.Context, req *pb.PaginationRequest) (*pb.Response, error) {
	pagination := utils.Pagination{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	}
	products, _, err := p.productUsecase.FindAllProducts(ctx, pagination)
	if err != nil {
		return &pb.Response{
			Message:    "failed to display products",
			Statuscode: 400,
			Errors:     "failed to display products",
		}, nil

	}
	fmt.Println(products)
	// Create a slice to hold the serialized protobuf messages
	serializedProducts := make([]*pb.ProductDetailsRequest, len(products))

	// Assuming products is a slice of domain.ProductDetails
	for i, product := range products {
		// Convert domain.ProductDetails to pb.ProductDetails
		protoProduct := &pb.ProductDetailsRequest{
			ProductName:    product.ProductName,
			ProductPrice:   int32(product.ProductPrice),
			ProductQuanity: int32(product.ProductQuantity),
			Category:       product.Category_name,
		}
		serializedProducts[i] = protoProduct
	}
	fmt.Println(serializedProducts)
	return &pb.Response{
		Message:    "successfully displayed products",
		Statuscode: 200,
		Errors:     "",
		Resp:       serializedProducts,
	}, nil
}
