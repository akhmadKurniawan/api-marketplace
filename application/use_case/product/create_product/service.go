package create_product

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"
	"log"
)

type CreateProductService struct {
	productRepository     infrastructure.ProductRepository
	shopRepository        infrastructure.ShopRepository
	productTypeRepository infrastructure.ProductTypeRepository
}

func NewCreateProductService(productRepo infrastructure.ProductRepository, shopRepo infrastructure.ShopRepository, productTypeRepo infrastructure.ProductTypeRepository) CreateProductService {
	return CreateProductService{
		productRepository:     productRepo,
		shopRepository:        shopRepo,
		productTypeRepository: productTypeRepo,
	}
}

func (s *CreateProductService) CreateProduct(ctx context.Context, req CreateProductRequest, img string) error {
	productType, err := s.productTypeRepository.GetProductTypeById(ctx, req.ProductType)
	if err != nil {
		log.Println("Service - CreateProduct err : ", err)
		return err
	}

	shop, err := s.shopRepository.GetShopById(ctx, req.ShopId)
	if err != nil {
		log.Println("Service - CreateProduct err : ", err)
		return err
	}

	product, errGetProduct := s.productRepository.GetProduct(ctx, models.Product{
		ShopId: req.ShopId,
		Name:   req.Name,
	})

	// kondisi eror = nil dan dilakukan pengecekan manual dan jika benar aplikasi berhenti
	if errGetProduct != nil || product.Name != "" {
		errGetProduct = errors.New("product name already exist")
		return errGetProduct
	}

	req.ProductType = productType.Model.ID
	req.ShopId = shop.Model.ID

	errCreate := s.productRepository.CreateProduct(ctx, RequestMapper(req), img)
	if errCreate != nil {
		log.Println("Service - CreateProduct error : ", errCreate)
		return errCreate
	}
	return nil
}
