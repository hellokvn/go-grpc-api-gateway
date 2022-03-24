package client

import (
	"fmt"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

func InitProductServiceClient() pb.ProductServiceClient {
	fmt.Println("Product Service Client")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
