package handler

import (
	"chapter3-sesi2/database"
	"chapter3-sesi2/repository/product_repository/product_pg"
	"chapter3-sesi2/repository/user_repository/user_pg"
	"chapter3-sesi2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()
	db := database.GetDatabaseInstance()
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)
	productRepo := product_pg.NewProductPG(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)
	authService := service.NewAuthService(userRepo, productRepo)

	route := gin.Default()
	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}
	productRoute := route.Group("/products")
	{
		productRoute.POST("/", authService.Aunthentication(), productHandler.CreateProduct)
		productRoute.PUT("/:productId", authService.Aunthentication(), authService.Authorization(), productHandler.UpdateProductById)
	}
	route.Run(":", port)
}
