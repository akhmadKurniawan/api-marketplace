package delete_user

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type DeleteUserHandler struct {
	userService DeleteUserService
}

func NewDeleteUserHandler(userServ DeleteUserService) DeleteUserHandler {
	return DeleteUserHandler{
		userService: userServ,
	}
}

func (h *DeleteUserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := h.userService.DeleteUser(ctx, id)
	if err != nil {
		log.Println("Controller - DeleteUser error while access service : ", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Success", true))
}
