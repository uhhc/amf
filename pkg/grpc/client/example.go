package client

import (
	"context"

	"github.com/uhhc/sdk-common-go/log"

	"github.com/uhhc/amf/pkg/grpc/pb"
)

// ExampleClient is the example client interface
type ExampleClient interface {
	Get(id string) (*pb.Example, error)
}

// ExampleGRPCClient is the example grpc client struct
type ExampleGRPCClient struct {
	client *GRPCClient
	logger log.Logger
}

// NewExampleGRPCClient to return a ExampleGRPCClient
func NewExampleGRPCClient(host, port string, logger log.Logger) ExampleClient {
	return &ExampleGRPCClient{
		client: NewGRPCClient(host, port, logger),
		logger: logger,
	}
}

// Get to get data from example grpc server
func (e *ExampleGRPCClient) Get(id string) (*pb.Example, error) {
	conn, err := e.client.OpenConn()
	if err != nil {
		e.logger.Errorw("connect to grpc server error", "error", err)
		return nil, err
	}
	defer conn.Close()

	e.logger.Infow("calling example grpc method: Get")
	client := pb.NewExampleServiceClient(conn)
	res, err := client.Get(context.Background(), &pb.GetRequest{Id: id})
	if err != nil {
		e.logger.Errorw("client get example error", "error", err)
		return nil, err
	}
	return res, nil
}
