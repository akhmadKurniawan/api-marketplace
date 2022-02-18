package create_product_type

import (
	"app/application/infrastructure"
	"context"
	"log"
)

type CreateProductTypeService struct {
	productTypeRepository infrastructure.ProductTypeRepository
}

func NewCreateProductTypeService(productTypeRepo infrastructure.ProductTypeRepository) CreateProductTypeService {
	return CreateProductTypeService{
		productTypeRepository: productTypeRepo,
	}
}

func (s *CreateProductTypeService) CreateProductType(ctx context.Context, req CreateProductTypeRequest, img string) error {
	errCreate := s.productTypeRepository.CreateProductType(ctx, RequestMapper(req, img))
	if errCreate != nil {
		log.Println("Service - CreateProductType error : ", errCreate)
		return errCreate
	}
	return nil
}
