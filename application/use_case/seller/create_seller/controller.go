package create_seller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateSellerHandler struct {
	sellerService CreateSellerService
}

func NewCreateSellerHandler(sellerServ CreateSellerService) CreateSellerHandler {
	return CreateSellerHandler{
		sellerService: sellerServ,
	}
}

func (h *CreateSellerHandler) CreateSeller(c *gin.Context) {
	req := CreateSellerRequest{}
	ctx := c.Request.Context()

	if err := c.Bind(&req); err != nil {
		log.Fatal("Controller - CreateSeller error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateSeller error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	errCreate := h.sellerService.CreateSeller(ctx, req)
	if errCreate != nil {
		log.Fatal("Controller - CreateSeller error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
