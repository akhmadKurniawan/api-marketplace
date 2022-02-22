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
	r, _ := c.Get("Role")
	accRole := strconv.FormatInt(r.(int64), 10)
	role, _ := strconv.Atoi(accRole)

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Controller - CreateSeller error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}
	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - CreateSeller error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if role < 2 {
		fmt.Println("you cannt create seller")
		c.JSON(400, response.SetMessage("you cannt create seller", false))
		return
	} else {
		errCreate := h.sellerService.CreateSeller(ctx, req, id)
		if errCreate != nil {
			log.Println("Controller - CreateSeller error while access service : ", errCreate)
			c.JSON(500, response.SetMessage(errCreate.Error(), false))
			return
		}
		c.JSON(http.StatusCreated, response.SetMessage("success", true))
	}
}
