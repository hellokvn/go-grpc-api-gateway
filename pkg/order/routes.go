package order

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine) {
	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/order")
	routes.Use(auth.AuthRequired)
	routes.POST("/", c.CreateOrder)
}

func (c *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, c.Client)
}
