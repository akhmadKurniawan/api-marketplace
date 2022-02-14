package delete_seller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type DeleteSellerHandler struct {
	sellerService DeleteSellerService
}

func NewDeleteSellerHandler(sellerServ DeleteSellerService) DeleteSellerHandler {
	return DeleteSellerHandler{
		sellerService: sellerServ,
	}
}

func (h *DeleteSellerHandler) DeleteSeller(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := h.sellerService.DeleteSeller(ctx, id)
	if err != nil {
		log.Fatal("Controller - DeleteSeller error while access service : ", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Success", true))
}
