package commons

import "errors"

// ErrorI interface describe an error based on own self error API response used in all platform APIs
//
// Reference implementation: ErrorS.
type ErrorI interface {
	ErrorStatus() int
	ErrorMessage() string
	ErrorDetails() map[string]any
	Error() string
	Unwrap() error
}

// ErrorCMap represent an ErrorI as should be added in the commons error context.
type ErrorCMap map[string]interface{}

// Unwrap implement an unwrapper error method in chaining based on ErrorI and ErrorCMap.
//
// This method is used internally for to generate a list of errors necessary in the commons error context.
//
// Note: this implementation use ErrorS as intermediate struct for generate ErrorCMap.
func Unwrap(err ErrorI) []ErrorCMap {
	var errs []ErrorI

	var errW error
	for errW = err; errW != nil; errW = errors.Unwrap(errW) {
		if errW.Error() == "" {
			continue
		}

		var errorCtx ErrorI
		if ok := errors.As(errW, &errorCtx); ok {
			errs = append(errs, errorCtx)
		} else {
			errs = append(errs, &ErrorS{err: errW})
		}
	}

	var errsCMap []ErrorCMap

	for _, errorI := range errs {
		e := ErrorCMap{}
		if errorI.ErrorStatus() > 0 {
			e["status"] = errorI.ErrorStatus()
		}
		if errorI.ErrorMessage() != "" {
			e["message"] = errorI.ErrorMessage()
		}
		if errorI.ErrorDetails() != nil {
			e["details"] = errorI.ErrorDetails()
		}
		errsCMap = append(errsCMap, e)
	}

	return errsCMap
}
