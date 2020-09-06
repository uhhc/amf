package example

import (
	"github.com/uhhc/amf/pkg/application/example/repository"
)

// CreateRequest holds the request parameters for the Create method.
type CreateRequest struct {
	Example repository.Example
}

// UpdateRequest holds the request parameters for the Update method.
type UpdateRequest struct {
	Example repository.Example
}

// ListRequest holds the request parameters for the List method.
type ListRequest struct {
	Where repository.Example `json:"where"`
	Order string             `json:"order"`
	Limit int32              `json:"limit"`
	Page  int32              `json:"page"`
}

// GetRequest holds the request parameters for the Get method.
type GetRequest struct {
	ID string `json:"id"`
}

// ChangeStatusRequest holds the request parameters for the ChangeStatus method.
type ChangeStatusRequest struct {
	ExampleID string `json:"example_id"`
	Status    string `json:"status"`
}

// DeleteRequest holds the request parameters for the Delete method.
type DeleteRequest struct {
	ID string `json:"id"`
}
