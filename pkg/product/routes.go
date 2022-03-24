package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/client"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/pb"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func RegisterProductRoutes(r *gin.Engine) {
	c := &ProductServiceClient{
		Client: client.InitProductServiceClient(),
	}

	routes := r.Group("/products")
	routes.POST("/", c.AddProduct)
	routes.GET("/:id", c.FindOne)
}

func (c *ProductServiceClient) FindOne(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (c *ProductServiceClient) AddProduct(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
