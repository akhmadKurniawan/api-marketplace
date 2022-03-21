package update_user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type UpdateUserHandler struct {
	userService UpdateUserService
}

func NewUpdateUserHandler(userServ UpdateUserService) UpdateUserHandler {
	return UpdateUserHandler{
		userService: userServ,
	}
}

func (h *UpdateUserHandler) UpdateUserHandler(c *gin.Context) {
	req := UpdateUserRequest{}
	id := c.Param("id")

	if err := c.ShouldBind(&req); err != nil {
		log.Println("Controller - Update error while binding request to json :", err)
		c.JSON(500, response.SetMessage(err.Error(), false))
		return
	}

	fmt.Println(req)
	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - Update validation :", err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(err.Error(), false))
		return
	}

	res, err := h.userService.UpdateUser(c.Request.Context(), req, id)
	if err != nil {
		log.Println("Controller - Update error while accessing service :", err)
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
	}

	c.JSON(http.StatusOK, SetResponse(res, "Success", true))

}
