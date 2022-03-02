package get_product_shopid

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowProductByShopIDResponse struct {
		Data ShowProductByShopIDResponseData `json:"data"`
		base.BaseResponse
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
)

func SetResponse(res *ShowProductByShopIDResponseData, message string, success bool) ShowProductByShopIDResponse {
	return ShowProductByShopIDResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(res),
	}
}

func ResponseMapper(model *ShowProductByShopIDResponseData) ShowProductByShopIDResponseData {
	return ShowProductByShopIDResponseData{
		ID:          model.ID,
		ProductType: model.ProductType,
		ShopID:      model.ShopID,
		Name:        model.Name,
		Price:       model.Price,
		Description: model.Description,
		Qty:         model.Qty,
		Image:       model.Image,
	}
}
