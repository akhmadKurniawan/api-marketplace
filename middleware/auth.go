package middleware

import (
	"app/app"
	"app/models"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type Token struct {
	UserID int64
	Name   string
	Role   int64
	jwt.StandardClaims
}

// AuthenticationRequired for auth
func AuthenticationRequired() gin.HandlerFunc {
	var (
		accToken models.UserToken
	)

	return func(c *gin.Context) {
		notAuth := []string{
			"/api/v1/users",
			"/api/v1/login",
			// "/api/v1/transactions",
		} //List of endpoints that doesn't require auth
		requestPath := c.Request.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath || strings.HasPrefix(requestPath, value) {
				c.Next()
				return
			}
		}

		tokenHeader := c.Request.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			c.JSON(403, response.SetMessage("Missing Auth Token", false))
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			c.JSON(403, response.SetMessage("Invalid/Malformed auth token", false))
			c.Abort()
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		c.Set("Role", tk.Role)

		if err != nil { //Malformed token, returns with http code 403 as usual
			c.JSON(403, response.SetMessage(fmt.Sprintf("Malformed authentication token: %s", err.Error()), false))
			c.Abort()
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			c.JSON(403, response.SetMessage("Token is not valid.", false))
			c.Abort()
			return
		}

		// Check token
		if err = app.GetDBP().Table("user_tokens").Select("token").Where(&models.UserToken{Token: tokenPart}).First(&accToken).Error; err != nil {
			c.JSON(403, response.SetMessage("Session login Anda berubah. Silakan lakukan login ulang", false))
			c.Abort()
			return
		}

		if tk.Role == 2 && strings.HasPrefix(requestPath, "/api/v1/admin") {
			c.JSON(403, response.SetMessage("You are not allowed acceess this route", false))
			c.Abort()
			return
		}

		c.Set("UserId", tk.UserID)
		c.Set("Name", tk.Name)
		c.Next() //proceed in the core chain!
	}
}
