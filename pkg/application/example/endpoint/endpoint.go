package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/uhhc/amf/pkg/application/example/endpoint/example"
	"github.com/uhhc/amf/pkg/application/example/service"
)

// Endpoints holds all Go kit endpoints for the Example service.
type Endpoints struct {
	Create       endpoint.Endpoint
	Update       endpoint.Endpoint
	List         endpoint.Endpoint
	Get          endpoint.Endpoint
	GetFromGRPC  endpoint.Endpoint
	ChangeStatus endpoint.Endpoint
	Delete       endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Example service.
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Create:       makeCreateEndpoint(s),
		Update:       makeUpdateEndpoint(s),
		List:         makeListEndpoint(s),
		Get:          makeGetEndpoint(s),
		GetFromGRPC:  makeGetFromGRPCEndpoint(s),
		ChangeStatus: makeChangeStatusEndpoint(s),
		Delete:       makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.CreateRequest) // type assertion
		err := s.Create(req.Example)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.UpdateRequest)
		err := s.Update(req.Example)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeListEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.ListRequest)
		ex, err := s.List(req.Where, req.Order, req.Limit, req.Page)
		if err != nil {
			return nil, err
		}
		return ex, nil
	}
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.GetRequest)
		ex, err := s.Get(req.ID)
		if err != nil {
			return nil, err
		}
		return ex, nil
	}
}

func makeGetFromGRPCEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.GetRequest)
		ex, err := s.GetFromGRPC(req.ID)
		if err != nil {
			return nil, err
		}
		return ex, nil
	}
}

func makeChangeStatusEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.ChangeStatusRequest)
		err := s.ChangeStatus(req.ExampleID, req.Status)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(example.DeleteRequest)
		err := s.Delete(req.ID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
