package get_product_shopid

import (
	"app/application/infrastructure"
	"context"
	"errors"
	"log"
)

type ShowProductByShopIDService struct {
	productRepository infrastructure.ProductRepository
}

func NewShowProductByShopIDService(productRepo infrastructure.ProductRepository) ShowProductByShopIDService {
	return ShowProductByShopIDService{
		productRepository: productRepo,
	}
}

func (s *ShowProductByShopIDService) ShowProductByShopID(ctx context.Context, id int) (*Response, error) {
	product, err := s.productRepository.GetProductByShopID(ctx, id)
	if err != nil {
		log.Println("ProductService - ShowProductByShopID error :", err)
		return nil, err
	}

	var data []ShowProductByShopIDResponseData
	for _, val := range product {
		res := ShowProductByShopIDResponseData{
			ShopID: val.ShopId,
		}
		data = append(data, res)
	}

	if err != nil || data == nil {
		err = errors.New("product not found")
		return nil, err
	}

	return &Response{Product: product}, nil
}
