// Package response contains the common responses in a client-server model
//
// NOTE: this package should be in a library
package response

type (
	// Common is the most common response to http requests
	// TODO optimize the serialization of this structure to avoid the use of reflection
	// TODO implement the interface json.Marshaler
	Common struct {
		// Code indicates the error code
		Code string `json:"code,omitempty"`
		// Message response message
		Message string `json:"message,omitempty"`
	}

	// Page is the common response for a paginated queries
	Page[T any] struct {
		Results      []T    `json:"results"`
		TotalPages   uint64 `json:"totalPages"`
		TotalResults uint64 `json:"totalResults"`
	}

	// InfiniteScroll standard response for
	InfiniteScroll[T any] struct {
		Results []T    `json:"results"`
		Cursor  string `json:"cursor"`
	}

	// Created common response for requests made to create a resource
	Created struct {
		ID any `json:"id,omitempty"`
	}
)
