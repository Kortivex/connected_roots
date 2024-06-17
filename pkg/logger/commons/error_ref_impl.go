package commons

import "errors"

// ErrorS is reference implementation of ErrorI.
//
// You can use this struct if yours needs are simples, although ideally you might want to
// implement your own implementation or extend from this as this is very basic for most cases.
//
// This implementation is reduced with the idea of being extended. For example
//
//		 type MyCustomError struct {
//			    ErrorS
//	         ...
//		 }
type ErrorS struct {
	Status  int            `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"details,omitempty"`
	err     error
}

// NewErrorS return a ErrorI instance based on ErrorS.
func NewErrorS(status int, message string, details map[string]any, err error) ErrorI {
	return &ErrorS{
		Status:  status,
		Message: message,
		Details: details,
		err:     err,
	}
}

// NewDefaultErrorS return a ErrorI instance based on ErrorS.
//
// This constructor use a default values for status, message or details. Is a shortcut for:
//
//	NewErrorS(0, "", nil, err)
func NewDefaultErrorS(err error) ErrorI {
	return NewErrorS(0, "", nil, err)
}

func (e *ErrorS) ErrorStatus() int {
	return e.Status
}

func (e *ErrorS) ErrorMessage() string {
	if e.Message == "" {
		return e.err.Error()
	}
	return e.Message
}

func (e *ErrorS) ErrorDetails() map[string]any {
	return e.Details
}

func (e *ErrorS) Error() string {
	return e.err.Error()
}

func (e *ErrorS) Unwrap() error {
	return errors.Unwrap(e.err)
}
