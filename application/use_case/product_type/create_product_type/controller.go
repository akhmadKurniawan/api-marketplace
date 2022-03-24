package create_product_type

import (
	"app/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	r, _ := c.Get("Role")
	accRole := strconv.FormatInt(r.(int64), 10)
	role, _ := strconv.Atoi(accRole)

	if err := c.Bind(&req); err != nil {
		log.Println("Controller - CreateProductType error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - CreateProductType error validation : ", err)
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

	if role < 3 {
		log.Println("Controller - CreateProductType: you cannt create product type")
		c.JSON(400, response.SetMessage("Controller - CreateProductType: you cannt create product type", false))
		return
	} else {
		errCreate := h.productTypeService.CreateProductType(ctx, req, file.FileUrl)
		if errCreate != nil {
			log.Println("Controller - CreateProductType error while access service : ", errCreate)
			c.JSON(500, response.SetMessage(errCreate.Error(), false))
			return
		}

		c.JSON(http.StatusCreated, response.SetMessage("success", true))
	}
}
