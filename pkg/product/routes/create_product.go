package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/pb"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	fmt.Println("Register")
	body := CreateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
