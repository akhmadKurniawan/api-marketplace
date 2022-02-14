package create_product

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateProductHandler struct {
	productService CreateProductService
}

func NewCreateProductHandler(productServ CreateProductService) CreateProductHandler {
	return CreateProductHandler{
		productService: productServ,
	}
}

func (h *CreateProductHandler) CreateProduct(c *gin.Context) {
	req := CreateProductRequest{}
	ctx := c.Request.Context()

	if err := c.ShouldBind(&req); err != nil {
		log.Fatal("Controller - CreateProduct error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateProduct error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	errCreate := h.productService.CreateProduct(ctx, req)
	if errCreate != nil {
		log.Fatal("Controller - CreateProduct error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
