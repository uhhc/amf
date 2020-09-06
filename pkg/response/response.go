package response

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/uhhc/sdk-common-go/log"
	"github.com/uhhc/sdk-common-go/types/constant/language"

	"github.com/uhhc/amf/pkg/config/errorcode"
)

// Header is the header field of CommonResponse
type Header struct {
	ErrNo  int32  `json:"err_no" example:"200"`
	ErrMsg string `json:"err_msg" example:"success"`
}

// CommonResponse is the common response for RESTFUL API
type CommonResponse struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

// NewCommonResponse return a CommonResponse instance
func NewCommonResponse(body interface{}, err *Error) *CommonResponse {
	if body == nil {
		body = struct{}{}
	}
	if err != nil {
		return &CommonResponse{
			Header: Header{
				ErrNo:  err.Code,
				ErrMsg: err.Error(),
			},
			Body: struct{}{},
		}
	}
	return &CommonResponse{
		Header: Header{
			ErrNo:  http.StatusOK,
			ErrMsg: "success",
		},
		Body: body,
	}
}

// FailureHeader for header when request failure
type FailureHeader struct {
	ErrNo  int32  `json:"err_no" example:"1003"`
	ErrMsg string `json:"err_msg" example:"请求参数错误"`
}

// ErrorResponse is the response for error
type ErrorResponse struct {
	Header FailureHeader `json:"header"`
	Body   struct{}      `json:"body"`
}

// Error is the struct to encapsule business error
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewError return an Error instance
func NewError(code int32, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// NewErrorFromCode to get error response from code
func NewErrorFromCode(code int32, args ...interface{}) *Error {
	// get language
	lang := viper.GetString("ERROR_LANG")
	if lang == "" {
		lang = "cn"
	}

	var msg string
	if len(args) != 0 {
		msg = fmt.Sprintf(errorcode.CodeMap[code][lang], args...)
	} else {
		msg = errorcode.CodeMap[code][lang]
	}

	return NewError(code, msg)
}

// Error implements the error interface
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// EncodeResponse encode the response to transport handler
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// EncodeGRPCResponse encode the GRPC response to transport handler
func EncodeGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// EncodeErrorResponse encode the error response to transport handler
func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	// Init logger
	logger := log.NewLogger()
	defer logger.FlushLogger()
	logger.Errorw("system error", "error", err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ErrorResponse{
		Header: FailureHeader{
			ErrNo:  errorcode.SystemError,
			ErrMsg: errorcode.CodeMap[errorcode.SystemError][language.CN],
		},
		Body: struct{}{},
	})
}
