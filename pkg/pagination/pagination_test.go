package pagination

import (
	"testing"
	"time"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/stretchr/testify/assert"
)

// Verify function returns a valid paginator when all parameters are valid.
func TestCreatePaginatorWithValidParams(t *testing.T) {
	// Given
	params := &PaginatorParams{
		Limit:          10,
		Order:          paginator.ASC,
		NextCursor:     "bextCursor==",
		PreviousCursor: "",
		Sort:           []string{"date"},
	}
	rules := []paginator.Rule{{Key: "date"}}

	// When
	result, err := CreatePaginator(params, rules)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

// Handle error when the next cursor is invalid base64 string.
func TestCreatePaginatorInvalidNextCursor(t *testing.T) {
	// Given
	params := &PaginatorParams{
		Limit:          10,
		Order:          paginator.ASC,
		NextCursor:     "invalid-base64",
		PreviousCursor: "",
		Sort:           []string{"user_id"},
	}
	rules := []paginator.Rule{{Key: "user_id"}}

	// When
	result, err := CreatePaginator(params, rules)

	// Then
	assert.Error(t, err)
	assert.Nil(t, result)
}

// Handle error when the previous cursor is invalid base64 string.
func TestCreatePaginatorInvalidPreviousCursor(t *testing.T) {
	// Given
	params := &PaginatorParams{
		Limit:          10,
		Order:          paginator.DESC,
		NextCursor:     "",
		PreviousCursor: "invalid-base64",
		Sort:           []string{"timestamp"},
	}
	rules := []paginator.Rule{{Key: "timestamp"}}

	// When
	result, err := CreatePaginator(params, rules)

	// Then
	assert.Error(t, err)
	assert.Nil(t, result)
}

// Assess performance implications of decoding and re-encoding cursors.
func TestCreatePaginatorPerformanceCursors(t *testing.T) {
	// Given
	params := &PaginatorParams{
		Limit:          10,
		Order:          paginator.ASC,
		NextCursor:     "bmV4dEN1cnNvcg==",
		PreviousCursor: "cHJldkN1cnNvcg==",
		Sort:           nil,
	}
	rules := []paginator.Rule{{Key: "id"}}

	start := time.Now()

	// When
	_, err := CreatePaginator(params, rules)

	duration := time.Since(start)

	// Then
	assert.NoError(t, err)
	assert.LessOrEqual(t, duration.Milliseconds(), int64(100), "Performance issue with cursor decoding/re-encoding")
}
