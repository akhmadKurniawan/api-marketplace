package create_seller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatInt(acc.(int64), 10)
	id, _ := strconv.Atoi(accountID)
	fmt.Println(id)
	a, _ := c.Get("Role")
	b := strconv.FormatInt(a.(int64), 10)
	d, _ := strconv.Atoi(b)
	fmt.Println("role", d)

	if err := c.ShouldBind(&req); err != nil {
		log.Fatal("Controller - CreateSeller error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}
	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateSeller error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}
	// userID = req.UserID
	errCreate := h.sellerService.CreateSeller(ctx, req, id)
	if errCreate != nil {
		log.Fatal("Controller - CreateSeller error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
