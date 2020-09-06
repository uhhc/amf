package grpc

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/uhhc/sdk-common-go/db"
	"github.com/uhhc/sdk-common-go/log"
	timeUtil "github.com/uhhc/sdk-common-go/util/time"

	"github.com/uhhc/amf/pkg/application/example/endpoint"
	"github.com/uhhc/amf/pkg/application/example/endpoint/example"
	"github.com/uhhc/amf/pkg/application/example/repository"
	"github.com/uhhc/amf/pkg/application/example/service"
	"github.com/uhhc/amf/pkg/grpc/pb"
	"github.com/uhhc/amf/pkg/response"
)

// Server is the GRPC Server struct
type Server struct {
	logger log.Logger
	get    gt.Handler
}

// NewServer to return a GRPC Server
func NewServer(_ context.Context, db *db.DbClient, logger log.Logger) pb.ExampleServiceServer {
	// Create Example Service and Endpoint
	repo := repository.NewRepository(db, logger)
	es := service.NewBusiness(repo, logger)
	ep := endpoint.MakeEndpoints(es)

	return &Server{
		logger: logger,
		get: gt.NewServer(
			ep.Get,
			DecodeGetExampleRequest,
			response.EncodeGRPCResponse,
		),
	}
}

// Get to get example data by grpc
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.Example, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		s.logger.Errorw("grpc method Get error", "resp", resp, "err", err)
		return nil, err
	}
	res := &pb.Example{
		Id:          resp.(*repository.Example).ID,
		ExampleId:   resp.(*repository.Example).ExampleID,
		ExampleName: resp.(*repository.Example).ExampleName,
		Status:      resp.(*repository.Example).Status,
		CreateTime:  timeUtil.GetFullTimeFormat(resp.(*repository.Example).CreateTime),
	}
	return res, nil
}

// DecodeGetExampleRequest to decode example request of get method
func DecodeGetExampleRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetRequest)
	return example.GetRequest{
		ID: req.Id,
	}, nil
}
