package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/order")
	routes.POST("/", c.CreateOrder)
}

func (c *ServiceClient) CreateOrder(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
