package create_product_type

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateProductTypeHandler struct {
	productTypeService CreateProductTypeService
}

func NewCreateProductTypeHandler(productTypeServ CreateProductTypeService) CreateProductTypeHandler {
	return CreateProductTypeHandler{
		productTypeService: productTypeServ,
	}
}

func (h *CreateProductTypeHandler) CreateProductType(c *gin.Context) {
	req := CreateProductTypeRequest{}
	ctx := c.Request.Context()

	if err := c.Bind(&req); err != nil {
		log.Fatal("Controller - CreateProductType error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateProductType error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	errCreate := h.productTypeService.CreateProductType(ctx, req)
	if errCreate != nil {
		log.Fatal("Controller - CreateProductType error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
