package errors

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/Kortivex/connected_roots/pkg/logger/commons"
)

type ErrorResponse struct {
	Err commons.ErrorS `json:"error"`
}

func (e *ErrorResponse) ErrorStatus() int {
	return e.Err.Status
}

func (e *ErrorResponse) ErrorMessage() string {
	return e.Err.Message
}

func (e *ErrorResponse) ErrorDetails() map[string]any {
	return e.Err.Details
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("ResponseError %d: %s", e.Err.Status, e.Err.Message)
}

func (e *ErrorResponse) Unwrap() error {
	return nil
}

// NewErrorResponse This function checks if an error exists in the errorsMap, and if so, returns its associated value,
// otherwise it returns an internal server error.
func NewErrorResponse(c echo.Context, err error) error {
	unwrapped := errors.Unwrap(err)
	if unwrapped != nil {
		for {
			auxUnwrap := errors.Unwrap(unwrapped)
			if auxUnwrap == nil {
				break
			}
			unwrapped = auxUnwrap
		}
		err = unwrapped
	}

	errResponse := &ErrorResponse{Err: matchError(err)}
	if err = c.JSON(errResponse.Err.Status, errResponse); err != nil {
		return err
	}

	return errResponse
}

func matchError(err error) commons.ErrorS {
	switch e := err.(type) {
	case *commons.ErrorS:
		return *e
	}

	if value, ok := errorAPIMap[err.Error()]; ok {
		return value
	}

	return commons.ErrorS{
		Status:  http.StatusInternalServerError,
		Message: ErrSomethingWentWrong.Error(),
	}
}
