package main

import (
	"fmt"
	"log"
	"os"

	database "app/app"
	"app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()

	v1 := r.Group("/api/v1")
	db := database.DBInit()

	v1.Use(middleware.AuthenticationRequired())

	UserRoutes(v1, db)
	LoginRoutes(v1, db)

	env := godotenv.Load()
	if env != nil {
		fmt.Println(env)
	}
	port := os.Getenv("APP_PORT")

	if port == "" {
		log.Fatal(fmt.Sprintf("PORT must be set [%s]", port))
	}

	r.Run(":" + port)
}

func UserRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateUserHandler(db)

	v1 := route.Group("/user/create")
	{
		v1.POST("", crHandler.CreateUser)
	}
}

func LoginRoutes(route *gin.RouterGroup, db *gorm.DB) {
	loginHandler := LoginHandler(db)

	v1 := route.Group("/login")
	{
		v1.POST("", loginHandler.Login)
	}
}
