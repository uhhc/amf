package client

import (
	"time"

	"github.com/spf13/viper"
	"github.com/uhhc/sdk-common-go/log"
	"google.golang.org/grpc"
)

// Client is the grpc client interface
type Client interface {
	OpenConn() (*grpc.ClientConn, error)
}

// GRPCClient is the grpc client struct
type GRPCClient struct {
	host   string
	port   string
	logger log.Logger
}

// NewGRPCClient to return a GRPCClient
func NewGRPCClient(host, port string, logger log.Logger) *GRPCClient {
	return &GRPCClient{
		host:   host,
		port:   port,
		logger: logger,
	}
}

// OpenConn to open a grpc connection
func (c *GRPCClient) OpenConn() (*grpc.ClientConn, error) {
	c.logger.Infow("prepare to dial a grpc connection")

	timeout := viper.GetDuration("GRPC_CONN_TIMEOUT")
	if timeout == 0 {
		timeout = 10
	}
	cc, err := grpc.Dial(c.host+":"+c.port, grpc.WithInsecure(), grpc.WithTimeout(timeout*time.Second))
	if err != nil {
		c.logger.Errorw("Failed to start gRPC connection", "error", err)
		return nil, err
	}
	return cc, nil
}
