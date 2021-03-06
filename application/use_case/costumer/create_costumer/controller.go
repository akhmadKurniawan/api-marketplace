package create_costumer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateCostumerHandler struct {
	costumerService CreateCostumerService
}

func NewCreateCostumerHandler(costumerServ CreateCostumerService) CreateCostumerHandler {
	return CreateCostumerHandler{
		costumerService: costumerServ,
	}
}

func (h *CreateCostumerHandler) CreateCostumer(c *gin.Context) {
	req := CreateCostumerRequest{}
	ctx := c.Request.Context()
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatInt(acc.(int64), 10)
	userID, _ := strconv.Atoi(accountID)

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Controller - CreateCostumer error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - CreateCostumer error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	req.UserID = userID
	errCreate := h.costumerService.CreateCostumer(ctx, req)
	if errCreate != nil {
		log.Println("Controller - CreateCostumer error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
