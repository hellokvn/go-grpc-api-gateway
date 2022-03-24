package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/routes"
)

func RegisterRoutes(r *gin.Engine) {
	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/auth")
	routes.POST("/register", c.Register)
	routes.POST("/login", c.Login)
}

func (c *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, c.Client)
}

func (c *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, c.Client)
}
