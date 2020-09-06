package middleware

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/uhhc/amf/pkg/response"
)

// HTTPResponse is the middleware for http response
var HTTPResponse endpoint.Middleware = func(e endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ex, err := e(ctx, request)
		if err != nil {
			return response.NewCommonResponse(ex, err.(*response.Error)), nil
		}
		return response.NewCommonResponse(ex, nil), nil
	}
}
