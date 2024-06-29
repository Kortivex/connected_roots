package errors

import (
	"fmt"
	"net/http"

	"github.com/Kortivex/connected_roots/pkg/utils"

	"github.com/labstack/echo/v4"

	"github.com/Kortivex/connected_roots/pkg/logger/commons"
)

type APIResponseError struct {
	Err commons.ErrorS `json:"error"`
}

func (e *APIResponseError) ErrorStatus() int {
	return e.Err.Status
}

func (e *APIResponseError) ErrorMessage() string {
	return e.Err.Message
}

func (e *APIResponseError) ErrorDetails() map[string]any {
	return e.Err.Details
}

func (e *APIResponseError) Error() string {
	return fmt.Sprintf("APIResponseError %d: %s", e.Err.Status, e.Err.Message)
}

func (e *APIResponseError) Unwrap() error {
	return &e.Err
}

func NewErrorResponse(c echo.Context, err error) error {
	errRes := &APIResponseError{Err: *matchError(fmt.Errorf("%w", err))}
	if err := c.JSON(errRes.Err.Status, errRes); err != nil {
		return err
	}

	return errRes
}

func matchError(err error) *commons.ErrorS {
	switch e := utils.UnwrapErr(err).(type) {
	default:
		if value, ok := errorAPIMap[e.Error()]; ok {
			return commons.NewErrorS(value.Status, value.Message, value.Details, err).(*commons.ErrorS)
		}

		return commons.NewErrorS(http.StatusInternalServerError, ErrSomethingWentWrong.Error(), nil, err).(*commons.ErrorS)
	}
}
