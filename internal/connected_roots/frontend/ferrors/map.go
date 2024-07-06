package ferrors

import (
	"errors"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"gorm.io/gorm"
	"net/http"
)

var (
	// ErrSomethingWentWrong is returned when an undetermined error occurred.
	ErrSomethingWentWrong = errors.New("something went wrong")
	// ErrBodyBadRequestWrongBody is returned when a request body has an incorrect format.
	ErrBodyBadRequestWrongBody = errors.New("malformed body")
	// ErrInvalidPayload is returned when a request body failed json schema validation.
	ErrInvalidPayload         = errors.New("invalid payload")
	ErrQueryParamInvalidValue = errors.New("error in query parameter value")
	ErrPathParamInvalidValue  = errors.New("error in path parameter value")
	ErrNotFound               = errors.New(gorm.ErrRecordNotFound.Error())
	ErrDuplicateKey           = errors.New(gorm.ErrDuplicatedKey.Error())
)

var errorAPIMap = map[string]commons.ErrorS{
	// General Errors
	ErrBodyBadRequestWrongBody.Error(): {Status: http.StatusBadRequest, Message: ErrBodyBadRequestWrongBody.Error()},
	ErrInvalidPayload.Error():          {Status: http.StatusBadRequest, Message: ErrInvalidPayload.Error()},
	ErrQueryParamInvalidValue.Error():  {Status: http.StatusBadRequest, Message: ErrQueryParamInvalidValue.Error()},
	ErrPathParamInvalidValue.Error():   {Status: http.StatusBadRequest, Message: ErrPathParamInvalidValue.Error()},
	ErrNotFound.Error():                {Status: http.StatusNotFound, Message: ErrNotFound.Error()},
	ErrDuplicateKey.Error():            {Status: http.StatusBadRequest, Message: ErrDuplicateKey.Error()},
}
