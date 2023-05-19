package order

import (
	"github.com/gin-gonic/gin"
	"github.com/nusaib44/go-grpc-api-gateway/pkg/auth"
	"github.com/nusaib44/go-grpc-api-gateway/pkg/config"
	"github.com/nusaib44/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	// a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
