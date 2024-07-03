package pagination

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Method returns a formatted string with status and message.
func TestErrorMethodReturnsFormattedString(t *testing.T) {
	// Given
	resp := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  404,
			Message: "Not Found",
		},
	}

	// When
	result := resp.Error()

	// Then
	expected := "(ResponsePaginationError) Error 404: Not Found"
	assert.Equal(t, expected, result)
}

// Method handles standard input without errors.
func TestErrorMethodHandlesStandardInput(t *testing.T) {
	// Given
	resp := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  200,
			Message: "OK",
		},
	}

	// When
	result := resp.Error()

	// Then
	expected := "(ResponsePaginationError) Error 200: OK"
	assert.Equal(t, expected, result)
}

// Method handles an empty message in ErrorPagination.
func TestErrorMethodHandlesEmptyMessage(t *testing.T) {
	// Given
	resp := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  200,
			Message: "",
		},
	}

	// When
	result := resp.Error()

	// Then
	expected := "(ResponsePaginationError) Error 200: "
	assert.Equal(t, expected, result)
}

// Method handles zero status in ErrorPagination.
func TestErrorMethodHandlesZeroStatus(t *testing.T) {
	// Given
	resp := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  0,
			Message: "No Status",
		},
	}

	// When
	result := resp.Error()

	// Then
	expected := "(ResponsePaginationError) Error 0: No Status"
	assert.Equal(t, expected, result)
}

// Method handles very large integers for status.
func TestErrorMethodHandlesVeryLargeIntegersForStatus(t *testing.T) {
	// Given
	largeStatus := int(^uint(0) >> 1) // Max int value
	resp := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  largeStatus,
			Message: "Large Status",
		},
	}

	// When
	result := resp.Error()

	// Then
	expected := fmt.Sprintf("(ResponsePaginationError) Error %d: Large Status", largeStatus)
	assert.Equal(t, expected, result)
}

// Method handles special characters and whitespace in a message.
func TestErrorMethodHandlesSpecialCharactersAndWhitespaceInMessage(t *testing.T) {
	// Given
	resp := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  200,
			Message: "Special characters! @# $%^&*()_+",
		},
	}

	// When
	result := resp.Error()

	// Then
	expected := "(ResponsePaginationError) Error 200: Special characters! @# $%^&*()_+"
	assert.Equal(t, expected, result)
}

// Method's output is dependent on the values of ErrorPagination fields.
func TestErrorMethodOutputDependentOnValues(t *testing.T) {
	// Given
	resp1 := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  404,
			Message: "Not Found",
		},
	}
	resp2 := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  500,
			Message: "Internal Server Error",
		},
	}

	// When
	result1 := resp1.Error()
	result2 := resp2.Error()

	// Then
	assert.NotEqual(t, result1, result2)
}

// Method does not modify the original ErrorPagination object.
func TestErrorMethodDoesNotModifyOriginalObject(t *testing.T) {
	// Given
	original := ResponsePaginationError{
		Err: ErrorPagination{
			Status:  200,
			Message: "OK",
		},
	}

	copy := original

	// When
	_ = original.Error()

	// Then
	assert.Equal(t, copy, original)
}
