package clients

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/auth_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/order_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/product_service"
	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	ProductService() 	product_service.ProductServiceClient
	CategoryService() 	category_service.CategoryServiceClient
	OrderService() 		order_service.OrderServiceClient
	AuthService() 		auth_service.AuthServiceClient
}

type grpcClients struct {
	productService  	product_service.ProductServiceClient
	categoryService 	category_service.CategoryServiceClient
	orderService    	order_service.OrderServiceClient
	authService 		auth_service.AuthServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManagerI, error) {
	connProductService, err := grpc.Dial(
		conf.ProductServiceGrpcHost+conf.ProductServiceGrpcPort,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connCategoryService, err := grpc.Dial(
		conf.CategoryServiceGrpcHost+conf.CategoryServiceGrpcPort,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	connOrderService, err := grpc.Dial(
		conf.OrderServiceGrpcHost+conf.OrderServiceGrpcPort,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	connAuthService, err := grpc.Dial(
		conf.AuthServiceGrpcHost+conf.AuthServiceGrpcPort,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		productService:  	product_service.NewProductServiceClient(connProductService),
		categoryService: 	category_service.NewCategoryServiceClient(connCategoryService),
		orderService:    	order_service.NewOrderServiceClient(connOrderService),
		authService: 		auth_service.NewAuthServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) ProductService() product_service.ProductServiceClient {
	return g.productService
}

func (g *grpcClients) CategoryService() category_service.CategoryServiceClient {
	return g.categoryService
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) AuthService() auth_service.AuthServiceClient {
	return g.authService 
}
