package main

import (
	"github.com/gin-gonic/gin"
	auth "github.com/hellokvn/go-grpc-api-gateway/pkg/auth"
	order "github.com/hellokvn/go-grpc-api-gateway/pkg/order"
	product "github.com/hellokvn/go-grpc-api-gateway/pkg/product"
)

func main() {
	r := gin.Default()

	auth.RegisterRoutes(r)
	product.RegisterRoutes(r)
	order.RegisterRoutes(r)

	r.Run(":3000")
}
