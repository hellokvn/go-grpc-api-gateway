package product

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine) {
	c := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/product")
	routes.Use(auth.AuthRequired)
	routes.POST("/", c.CreateProduct)
	routes.GET("/:id", c.FindOne)
}

func (c *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, c.Client)
}

func (c *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, c.Client)
}
