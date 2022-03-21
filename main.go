package main

import (
	"fmt"
	"log"
	"os"

	database "app/app"
	docs "app/docs"
	"app/middleware"
	"app/shared"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()

	v1 := r.Group("/api/v1")
	db := database.DBInit()

	shared.InitLogger()

	ImgRoute(v1, db)
	v1.Use(middleware.AuthenticationRequired())

	UserRoutes(v1, db)
	LoginRoutes(v1, db)
	SellerRoutes(v1, db)
	CostumerRoutes(v1, db)
	ProductTypeRoutes(v1, db)
	ProductRoutes(v1, db)
	ShopRoutes(v1, db)
	TransactionRoutes(v1, db)
	WaletRoutes(v1, db)
	UpdateScheduler(db)

	env := godotenv.Load()
	if env != nil {
		fmt.Println(env)
	}
	port := os.Getenv("APP_PORT")

	if port == "" {
		log.Fatal(fmt.Sprintf("PORT must be set [%s]", port))
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + port)
}

func UserRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateUserHandler(db)
	dlHandler := DeleteUserHandler(db)
	upHandler := UpdateUserHandler(db)
	verifyHandler := VerifyEmailUserHandler(db)

	v1 := route.Group("/users")
	{
		v1.POST("", crHandler.CreateUser)
		v1.PUT("/:id", upHandler.UpdateUserHandler)
		v1.DELETE("/:id", dlHandler.DeleteUser)
		v1.GET("/active/:param", verifyHandler.VerifyEmailUser)
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
	getShopIdHandler := ShowProductByShopIDHandler(db)
	v1 := route.Group("/products")
	{
		v1.POST("", crHandler.CreateProduct)
		v1.GET("/:id", getShopIdHandler.ShowProductByShopID)
	}
}

func ImgRoute(route *gin.RouterGroup, db *gorm.DB) {
	route.GET("/images/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.File("./images/default-image.png")
		} else {
			c.File("./images/" + id)
		}
	})
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
	upHandler := UpdateTransactionHandler(db)
	getHandler := ShowTransactionHandler(db)

	v1 := route.Group("/transactions")
	{
		v1.GET("", getHandler.ShowTransaction)
		v1.POST("", crHandler.CreateTransaction)
		v1.POST("/:id", upHandler.UpdateTransaction)
	}
}

func WaletRoutes(route *gin.RouterGroup, db *gorm.DB) {
	crHandler := CreateWaletHandler(db)

	v1 := route.Group("/walets")
	{
		v1.POST("", crHandler.CreateWalet)
	}
}

func UpdateScheduler(db *gorm.DB) {
	fmt.Println("in")
	scheduler := cron.New()
	defer scheduler.Stop()

	upHandler := UpdateSchedulerHandler(db)
	// sch, err := scheduler.AddFunc("*/1 * * * *", upHandler.UpdateScheduler)
	// fmt.Println("s", sch, err)
	scheduler.AddFunc("*/60 * * * *", upHandler.UpdateScheduler)
	go scheduler.Start()

}
