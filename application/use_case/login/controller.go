package login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type LoginHandler struct {
	loginService LoginService
}

func NewLoginHandler(loginService LoginService) LoginHandler {
	return LoginHandler{
		loginService: loginService,
	}
}

// @BasePath /api/v1

// Login is endpoint for login to app
// @Summary Login To App
// @Description for login to app
// @Tags Login user
// @Accept json
// @Param request body LoginRequestBody{} true "Login User"
// @Produce json
// // @Failure 400 {object} models.BaseResponse
// //@Failure 404 {object} LoginResponse
// //@Failure 500 {object} LoginResponse
// @Success 200 {object} LoginResponse
// @Router /login [post]
func (h *LoginHandler) Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.Bind(&req); err != nil {
		log.Println("Controller - Login error while binding request to json :", err)
		c.JSON(400, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&req); !ok {
		log.Println("Controller - Login validation :", err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(err.Error(), false))
		return
	}

	res, err := h.loginService.LoginUser(c.Request.Context(), req)
	if err != nil {
		log.Println("Controller - Login error while accessing service :", err)
		c.JSON(400, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(res, "Success", true))
}
