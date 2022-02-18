package create_product_type

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"
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
	productType, err := s.productTypeRepository.GetProductType(ctx, models.ProductType{
		Name: req.Name,
	})
	if err != nil || productType.Name != "" {
		errGetProductType := errors.New("product type name already exist")
		return errGetProductType
	}

	errCreate := s.productTypeRepository.CreateProductType(ctx, RequestMapper(req, img))
	if errCreate != nil {
		log.Println("Service - CreateProductType error : ", errCreate)
		return errCreate
	}
	return nil
}
