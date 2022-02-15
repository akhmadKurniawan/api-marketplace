package create_walet

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateWaletHandler struct {
	waletService CreateWaletService
}

func NewCreateWaletHandler(waletServ CreateWaletService) CreateWaletHandler {
	return CreateWaletHandler{
		waletService: waletServ,
	}
}

func (h *CreateWaletHandler) CreateWalet(c *gin.Context) {
	req := CreateWaletRequest{}
	ctx := c.Request.Context()
	acc, _ := c.Get("UserId")
	accUser := strconv.FormatInt(acc.(int64), 10)
	userID, _ := strconv.Atoi(accUser)

	if err := c.ShouldBind(&req); err != nil {
		log.Fatal("Controller - CreateWalet error while binding request : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Fatal("Controller - CreateWalet error validation : ", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	errCreate := h.waletService.CreateWalet(ctx, req, userID)
	if errCreate != nil {
		log.Fatal("Controller - CreateWalet error while access service : ", errCreate)
		c.JSON(500, response.SetMessage(errCreate.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, response.SetMessage("success", true))
}
