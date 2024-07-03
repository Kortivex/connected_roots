package pagination

import "fmt"

const (
	ErrMsgInvalidSort    = "invalid sort value"
	ErrMsgInvalidCursor  = "invalid cursor value"
	ErrMsgInternalServer = "internal server error"
)

type ErrorPagination struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (e ErrorPagination) Error() string {
	return fmt.Sprintf("(ErrorPagination) Error %d: %s", e.Status, e.Message)
}

// ResponsePaginationError defines the structure of the error response used by the
// SmartWorks for PostgreSQL pagination.
type ResponsePaginationError struct {
	Err ErrorPagination `json:"error"`
}

func (e ResponsePaginationError) Error() string {
	return fmt.Sprintf("(ResponsePaginationError) Error %d: %s", e.Err.Status, e.Err.Message)
}
