package create_product

import (
	"app/middleware"
	"log"
	"net/http"
	"strconv"
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
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatInt(acc.(int64), 10)
	userID, _ := strconv.Atoi(accountID)

	// fmt.Println(req.Image)
	// //input tipe file
	// file, err := c.FormFile(req.Image.Filename)
	// fmt.Println(file)
	// fmt.Println(err)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
	// 	return
	// }
	// fmt.Println(req.Image.Filename)

	// //set folder untuk menyimpan filenya
	// path := "images/" + file.Filename
	// if err := c.SaveUploadedFile(file, path); err != nil {
	// 	c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
	// 	return
	// }

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

	file, errFile := middleware.UploadFile(c, "image")
	if errFile != nil {
		if !strings.Contains(errFile.Error(), "Empty File") {
			c.JSON(422, response.SetMessage(errFile.Error(), false))
			return
		}
	}

	req.UserID = userID
	errCreate := h.productService.CreateProduct(ctx, req, file.FileUrl)
	if errCreate != nil {
		log.Fatal("Controller - CreateProduct error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
