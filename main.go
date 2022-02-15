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
	SellerRoutes(v1, db)
	CostumerRoutes(v1, db)
	ProductTypeRoutes(v1, db)
	ProductRoutes(v1, db)
	ShopRoutes(v1, db)
	TransactionRoutes(v1, db)

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
	dlHandler := DeleteUserHandler(db)

	v1 := route.Group("/users")
	{
		v1.POST("", crHandler.CreateUser)
		v1.DELETE("/:id", dlHandler.DeleteUser)
	}
}

func CostumerRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateCostumerHandler(db)

	v1 := route.Group("/costumers")
	{
		v1.POST("", crHandler.CreateCostumer)
	}
}

func SellerRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateSellerHandler(db)
	dlHandler := DeleteSellerHandler(db)

	v1 := route.Group("/sellers")
	{
		v1.POST("", crHandler.CreateSeller)
		v1.DELETE("/:id", dlHandler.DeleteSeller)
	}
}

func LoginRoutes(route *gin.RouterGroup, db *gorm.DB) {
	loginHandler := LoginHandler(db)

	v1 := route.Group("/login")
	{
		v1.POST("", loginHandler.Login)
	}
}

func ProductRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateProductHandler(db)

	v1 := route.Group("/products")
	{
		v1.POST("", crHandler.CreateProduct)
	}
}

func ProductTypeRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateProductTypeHandler(db)

	v1 := route.Group("/product_types")
	{
		v1.POST("", crHandler.CreateProductType)
	}
}

func ShopRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateShopHandler(db)

	v1 := route.Group("/shops")
	{
		v1.POST("", crHandler.CreateShop)
	}
}

func TransactionRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateTransactionHandler(db)

	v1 := route.Group("/transactions")
	{
		v1.POST("", crHandler.CreateTransaction)
	}
}
