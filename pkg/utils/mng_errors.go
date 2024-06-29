package utils

import (
	"errors"
	"strings"
)

func UnwrapErr(err error) error {
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
	return err
}

func CheckErrorMsg(err error, msg string) bool {
	err = UnwrapErr(err)
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), msg)
}
