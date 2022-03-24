package product

import (
	"fmt"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	fmt.Println("Product Service Client")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
