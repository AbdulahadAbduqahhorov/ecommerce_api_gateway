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
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	
	docs.SwaggerInfo.Title = "Ecommerce API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use()

	handler := handlers.New(&handlers.HandlerOptions{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	//**Product
	router.POST("/product", handler.CreateProduct)
	router.GET("/product", handler.GetAllProduct)
	router.GET("/product/:product_id", handler.GetProductById)
	router.PUT("/product/:product_id", handler.UpdateProduct)
	router.DELETE("/product/:product_id", handler.DeleteProduct)

	//**Category
	router.POST("/category", handler.CreateCategory)
	router.GET("/category", handler.GetAllCategory)
	router.GET("/category/:category_id", handler.GetCategoryById)
	router.PUT("/category/:category_id", handler.UpdateCategory)
	router.DELETE("/category/:category_id", handler.DeleteCategory)

	

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
