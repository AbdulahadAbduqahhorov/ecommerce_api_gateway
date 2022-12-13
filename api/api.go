package api

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/docs"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/handlers"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/clients"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services clients.ServiceManagerI
}

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func New(opt *RouterOptions) *gin.Engine {

	docs.SwaggerInfo.Title = "Ecommerce API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(MyCORSMiddleware())
	handler := handlers.New(&handlers.HandlerOptions{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	//**Product
	router.POST("/product",handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}) ,handler.CreateProduct)
	router.GET("/product", handler.GetAllProduct)
	router.GET("/product/:product_id", handler.GetProductById)
	router.PUT("/product/:product_id", handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}),handler.UpdateProduct)
	router.DELETE("/product/:product_id", handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}),handler.DeleteProduct)

	//**Category
	category := router.Group("/category")
	{
		category.POST("",handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}), handler.CreateCategory)
		category.GET("",handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}), handler.GetAllCategory)
		category.GET("/:category_id", handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}),handler.GetCategoryById)
		category.PUT("/:category_id", handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}),handler.UpdateCategory)
		category.DELETE("/:category_id", handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}),handler.DeleteCategory)
	}

	//**Order
	router.POST("/order", handler.AuthMiddleware(map[string]bool{"*":true}),handler.CreateOrder)
	router.GET("/order", handler.AuthMiddleware(map[string]bool{"*":true}),handler.GetAllOrder)
	router.GET("/order/:order_id", handler.AuthMiddleware(map[string]bool{"*":true}),handler.GetOrderById)

	//**User
	router.POST("/user",handler.AuthMiddleware(map[string]bool{"SUPERADMIN":true}), handler.CreateUser)
	router.GET("/user", handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}),handler.GetAllUser)
	router.GET("/user/:user_id",handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}) ,handler.GetUserById)
	router.PUT("/user",handler.AuthMiddleware(map[string]bool{"ADMIN":true,"SUPERADMIN":true}), handler.UpdateUser)
	router.DELETE("/user/:user_id",handler.AuthMiddleware(map[string]bool{"SUPERADMIN":true}),handler.DeleteUser)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

// MyCORSMiddleware ...
func MyCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
