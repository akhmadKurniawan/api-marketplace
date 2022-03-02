package get_product_shopid

import (
	"app/models"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowProductByShopIDResponse struct {
		base.BaseResponse
		Data []ShowProductByShopIDResponseData `json:"data"`
	}
	ShowProductByShopIDResponseData struct {
		ID          int    `json:"id"`
		ProductType int    `json:"product_type"`
		ShopID      int    `json:"shop_id"`
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		Qty         int    `json:"qty"`
		Image       string `json:"image"`
	}
	Response struct {
		Product []models.Product
	}
)

func SetResponse(res *Response, message string, success bool) ShowProductByShopIDResponse {
	return ShowProductByShopIDResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMappers(res),
	}
}

func ResponseMappers(res *Response) []ShowProductByShopIDResponseData {
	var list []ShowProductByShopIDResponseData
	for _, val := range res.Product {
		response := ShowProductByShopIDResponseData{
			ID:          val.ID,
			ProductType: val.ProductType,
			ShopID:      val.ShopId,
			Name:        val.Name,
			Price:       val.Price,
			Description: val.Description,
			Qty:         val.Qty,
			Image:       val.Image,
		}
		list = append(list, response)
	}
	return list
}
