package create_product

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateProductService struct {
	productRepository infrastructure.ProductRepository
}

func NewCreateProductService(productRepo infrastructure.ProductRepository) CreateProductService {
	return CreateProductService{
		productRepository: productRepo,
	}
}

func (s *CreateProductService) CreateProduct(ctx context.Context, req CreateProductRequest) error {
	errCreate := s.productRepository.CreateProduct(ctx, RequestMapper(req))
	if errCreate != nil {
		log.Fatal("Service - CreateProduct error : ", errCreate)
	}
	return nil
}
