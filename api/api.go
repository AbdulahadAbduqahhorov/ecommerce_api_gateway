package api

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/handlers"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/clients"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @contact.url   http://example.com
// @contact.email example@swagger.io
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services clients.ServiceManagerI
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use()

	handler := handlers.New(&handlers.HandlerOptions{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})



	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
