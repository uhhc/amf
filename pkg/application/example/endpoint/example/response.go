package example

import (
	"github.com/uhhc/amf/pkg/application/example/repository"
	"github.com/uhhc/amf/pkg/response"
)

// CreateResponse represents the response of creating data
type CreateResponse struct {
	Header response.Header `json:"header"`
	Body   struct{}        `json:"body"`
}

// UpdateResponse represents the response of updating data
type UpdateResponse struct {
	Header response.Header `json:"header"`
	Body   struct{}        `json:"body"`
}

// ListResponse represents the response of listing data
type ListResponse struct {
	Header response.Header        `json:"header"`
	Body   repository.ExampleList `json:"body"`
}

// GetResponse represents the response of getting data by ID
type GetResponse struct {
	Header response.Header    `json:"header"`
	Body   repository.Example `json:"body"`
}

// ChangeStatusResponse represents the response of changing data status
type ChangeStatusResponse struct {
	Header response.Header `json:"header"`
	Body   struct{}        `json:"body"`
}

// DeleteResponse represents the response of deleting data
type DeleteResponse struct {
	Header response.Header `json:"header"`
	Body   struct{}        `json:"body"`
}
