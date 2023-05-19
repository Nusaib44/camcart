package order

import (
	"fmt"

	"github.com/nusaib44/go-grpc-api-gateway/pkg/config"
	"github.com/nusaib44/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {

	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	fmt.Println("connected to order server:", err)
	return pb.NewOrderServiceClient(cc)
}
