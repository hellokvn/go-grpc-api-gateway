package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/auth")
	routes.POST("/", c.Register)
	routes.GET("/:id", c.Login)
}

func (c *ServiceClient) Register(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *ServiceClient) Login(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
