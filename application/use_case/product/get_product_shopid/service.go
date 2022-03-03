package get_product_shopid

import (
	"app/application/infrastructure"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
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

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, "one", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(ctx, "one").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("redis", val)

	return &Response{Product: product}, nil
}
