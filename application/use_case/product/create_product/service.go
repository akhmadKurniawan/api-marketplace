package create_product

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateProductService struct {
	productRepository infrastructure.ProductRepository
	shopRepository    infrastructure.ShopRepository
	sellerRepository  infrastructure.SellerRepository
}

func NewCreateProductService(productRepo infrastructure.ProductRepository, shopRepo infrastructure.ShopRepository, sellerRepo infrastructure.SellerRepository) CreateProductService {
	return CreateProductService{
		productRepository: productRepo,
		shopRepository:    shopRepo,
		sellerRepository:  sellerRepo,
	}
}

func (s *CreateProductService) CreateProduct(ctx context.Context, req CreateProductRequest, img string) error {
	errCreate := s.productRepository.CreateProduct(ctx, RequestMapper(req), img)
	if errCreate != nil {
		log.Fatal("Service - CreateProduct error : ", errCreate)
	}
	return nil
}
