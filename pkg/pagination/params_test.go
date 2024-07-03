package pagination

import (
	"testing"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/stretchr/testify/assert"
)

// Function handles valid next_cursor and previous_cursor by decoding and re-encoding them.
func TestHandleValidCursors(t *testing.T) {
	// Given
	validCursor := "dGVzdA==" // base64 URL encoding of "test"
	params := PaginatorParams{
		NextCursor:     validCursor,
		PreviousCursor: validCursor,
	}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "dGVzdA==", params.NextCursor)
	assert.Equal(t, "dGVzdA==", params.PreviousCursor)
}

// Function sets default limit when limit is not specified or is zero.
func TestDefaultLimitSetting(t *testing.T) {
	// Given
	params := PaginatorParams{Limit: 0}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, defaultLimit, params.Limit)
}

// Function respects the limit when it is within the acceptable range.
func TestRespectWithinRangeLimit(t *testing.T) {
	// Given
	withinRangeLimit := 500
	params := PaginatorParams{Limit: withinRangeLimit}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, withinRangeLimit, params.Limit)
}

// Function sets default order when order is not specified.
func TestDefaultOrderSetting(t *testing.T) {
	// Given
	params := PaginatorParams{}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, defaultOrder, params.Order)
}

// Function maintains specified order when it is either ASC or DESC.
func TestMaintainSpecifiedOrder(t *testing.T) {
	// Given
	params := PaginatorParams{Order: paginator.DESC}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, paginator.DESC, params.Order)
}

// Function returns an error when next_cursor is not base64 URL encoded.
func TestErrorOnInvalidNextCursorEncoding(t *testing.T) {
	// Given
	invalidCursor := "not_base64_encoded"
	params := PaginatorParams{NextCursor: invalidCursor}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), ErrMsgInvalidCursor)
}

// Function returns an error when previous_cursor is not base64 URL encoded.
func TestErrorOnInvalidPreviousCursorEncoding(t *testing.T) {
	// Given
	invalidCursor := "not_base64_encoded"
	params := PaginatorParams{PreviousCursor: invalidCursor}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), ErrMsgInvalidCursor)
}

// Function caps the limit to defaultMaxLimit when the specified limit exceeds the maximum allowed.
func TestCapLimitToMaxAllowed(t *testing.T) {
	// Given
	excessiveLimit := 2000
	params := PaginatorParams{Limit: excessiveLimit}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, defaultMaxLimit, params.Limit)
}

// Function defaults to ASC order when an invalid order is specified.
func TestDefaultToAscOnInvalidOrder(t *testing.T) {
	// Given
	params := PaginatorParams{Order: paginator.Order("invalid")}

	// When
	err := getPaginatorParams(&params)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, paginator.ASC, params.Order)
}
