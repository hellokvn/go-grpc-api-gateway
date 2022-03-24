package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product"
)

func main() {
	r := gin.Default()

	product.RegisterProductRoutes(r)

	r.Run(":3000")
}
