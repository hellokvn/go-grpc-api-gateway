package order

import (
	"fmt"

	"github.com/hellokvn/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient() pb.OrderServiceClient {
	fmt.Println("Order Service Client")

	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
