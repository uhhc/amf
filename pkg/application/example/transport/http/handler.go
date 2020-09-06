package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/uhhc/sdk-common-go/db"
	"github.com/uhhc/sdk-common-go/log"

	"github.com/uhhc/amf/pkg/application/example/endpoint"
	"github.com/uhhc/amf/pkg/application/example/endpoint/example"
	"github.com/uhhc/amf/pkg/application/example/repository"
	"github.com/uhhc/amf/pkg/application/example/service"
	"github.com/uhhc/amf/pkg/middleware"
	"github.com/uhhc/amf/pkg/response"
)

var (
	// ErrBadRouting is an error for bad routing
	ErrBadRouting = errors.New("bad routing")
)

// MakeHTTPHandler make a http handler for http server
func MakeHTTPHandler(db *db.DbClient, logger log.Logger) http.Handler {
	// Create Example Service
	var es service.Service
	repo := repository.NewRepository(db, logger)
	es = service.NewBusiness(repo, logger)

	// Create Go kit endpoints for the Example Service
	var serverOptions []kithttp.ServerOption
	return NewHandler(endpoint.MakeEndpoints(es), serverOptions, logger)
}

// NewHandler wires Go kit endpoints to the HTTP transport.
func NewHandler(
	exEndpoints endpoint.Endpoints, options []kithttp.ServerOption, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints
	var (
		r            = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger))
		errorEncoder = kithttp.ServerErrorEncoder(response.EncodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)
	rd := NewRequestDecoder(&logger)

	r.Methods("POST").Path("/v1/examples").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.Create),
		rd.decodeCreateRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("GET").Path("/v1/examples").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.List),
		rd.decodeListRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("GET").Path("/v1/examples/{example_id}").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.Get),
		rd.decodeGetRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("GET").Path("/v1/examples/{example_id}/grpc").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.GetFromGRPC),
		rd.decodeGetRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/v1/examples/{example_id}").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.Update),
		rd.decodeUpdateRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/v1/examples/{example_id}/status").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.ChangeStatus),
		rd.decodeChangeStausRequest,
		response.EncodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/v1/examples/{example_ids}").Handler(kithttp.NewServer(
		middleware.HTTPResponse(exEndpoints.Delete),
		rd.decodeDeleteRequest,
		response.EncodeResponse,
		options...,
	))

	return r
}

// RequestDecoder represents the decoder for request
type RequestDecoder struct {
	logger *log.Logger
}

// NewRequestDecoder to return RequestDecoder instance
func NewRequestDecoder(logger *log.Logger) *RequestDecoder {
	return &RequestDecoder{
		logger: logger,
	}
}

// @Summary 创建一个样例
// @Description 该接口提供创建样例的API
// @Tags 样例
// @Accept json
// @Produce json
// @Param example body repository.Example true "要创建的样例的具体信息。通过 request body 发送"
// @Success 200 {object} example.CreateResponse "创建成功"
// @Failure 500 {object} response.ErrorResponse "创建失败"
// @Router /v1/examples [post]
func (rd *RequestDecoder) decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req example.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Example); e != nil {
		rd.logger.Errorw("json decode error", "error", e)
		return nil, e
	}
	return req, nil
}

// @Summary 更新一个样例
// @Description 该接口提供更新样例的API
// @Tags 样例
// @Accept json
// @Produce json
// @Param example_id path string true "样例ID。通过 URL Path 发送"
// @Param example body repository.Example true "要更新的样例的具体信息。通过 request body 发送"
// @Success 200 {object} example.UpdateResponse "更新成功"
// @Failure 500 {object} response.ErrorResponse "更新失败"
// @Router /v1/examples/{example_id} [put]
func (rd *RequestDecoder) decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req example.UpdateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Example); e != nil {
		rd.logger.Errorw("json decode error", "error", e)
		return nil, e
	}

	vars := mux.Vars(r)
	id, ok := vars["example_id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.Example.ExampleID = id

	return req, nil
}

// @Summary 获取某个样例详情
// @Description 该接口提供查看某个样例详情的API
// @Tags 样例
// @Accept json
// @Produce json
// @Param example_id path string true "样例ID。通过 URL Path 发送"
// @Success 200 {object} example.GetResponse "获取成功"
// @Failure 500 {object} response.ErrorResponse "获取失败"
// @Router /v1/examples/{example_id} [get]
func (rd *RequestDecoder) decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["example_id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return example.GetRequest{ID: id}, nil
}

// @Summary 获取样例列表
// @Description 该接口提供样例列表的API
// @Tags 样例
// @Accept json
// @Produce json
// @Param example_name query string false "样例名称"
// @Param page_size query int false "每页展示数量"
// @Param page_num query int false "第几页"
// @Success 200 {object} example.ListResponse "获取列表成功"
// @Failure 500 {object} response.ErrorResponse "获取失败"
// @Router /v1/examples [get]
func (rd *RequestDecoder) decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	e := r.ParseForm()
	if e != nil {
		rd.logger.Errorw("parse form error", "error", e)
		return nil, e
	}

	limit, _ := strconv.Atoi(r.Form.Get("page_size"))
	page, _ := strconv.Atoi(r.Form.Get("page_num"))

	req := example.ListRequest{
		Where: repository.Example{
			ExampleName: r.Form.Get("example_name"),
			Status:      r.Form.Get("status"),
		},
		Order: r.Form.Get("order"),
		Limit: int32(limit),
		Page:  int32(page),
	}

	return req, nil
}

// @Summary 更新某个样例状态
// @Description 该接口提供更新某个样例状态的API
// @Tags 样例
// @Accept json
// @Produce json
// @Param example_id path string true "样例ID。通过 URL Path 发送"
// @Param status body string true "要修改的样例的状态。通过 request body 发送"
// @Success 200 {object} example.ChangeStatusResponse "更新成功"
// @Failure 500 {object} response.ErrorResponse "更新失败"
// @Router /v1/examples/{example_id}/status [put]
func (rd *RequestDecoder) decodeChangeStausRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req example.ChangeStatusRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		rd.logger.Errorw("json decode error", "error", e)
		return nil, e
	}

	vars := mux.Vars(r)
	id, ok := vars["example_id"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.ExampleID = id

	return req, nil
}

// @Summary 删除一个或多个样例
// @Description 该接口提供删除一个或多个样例的API
// @Tags 样例
// @Accept json
// @Produce json
// @Param example_ids path string true "样例ID或ID列表，用英文逗号分隔。通过 URL Path 发送"
// @Success 200 {object} example.DeleteResponse "删除成功"
// @Failure 500 {object} response.ErrorResponse "删除失败"
// @Router /v1/examples/{example_ids} [delete]
func (rd *RequestDecoder) decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["example_ids"]
	if !ok {
		return nil, ErrBadRouting
	}
	return example.DeleteRequest{ID: id}, nil
}
