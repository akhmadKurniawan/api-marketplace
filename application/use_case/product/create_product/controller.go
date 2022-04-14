package create_product

import (
	"app/middleware"
	"log"
	"net/http"
	"strings"

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
		log.Println("Controller - CreateProduct error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - CreateProduct error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	file, errFile := middleware.UploadFile(c, "image")
	if errFile != nil {
		if !strings.Contains(errFile.Error(), "Empty File") {
			c.JSON(422, response.SetMessage(errFile.Error(), false))
			return
		}
	}

	errCreate := h.productService.CreateProduct(ctx, req, file.FileUrl)
	if errCreate != nil {
		log.Println("Controller - CreateProduct error while access service : ", errCreate)
		c.JSON(400, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
