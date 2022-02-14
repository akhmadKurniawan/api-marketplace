package create_shop

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateShopHandler struct {
	shopService CreateShopService
}

func NewCreateShopHandler(shopServ CreateShopService) CreateShopHandler {
	return CreateShopHandler{
		shopService: shopServ,
	}
}

func (h *CreateShopHandler) CreateShop(c *gin.Context) {
	req := CreateShopRequest{}
	ctx := c.Request.Context()

	if err := c.ShouldBind(&req); err != nil {
		log.Fatal("Controller - CreateShop error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateShop error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	errCreate := h.shopService.CreateShop(ctx, req)
	if errCreate != nil {
		log.Fatal("Controller - CreateShop error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
