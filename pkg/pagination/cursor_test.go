package pagination

import (
	"encoding/base64"
	"testing"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/stretchr/testify/assert"
)

// Encode valid 'Before' cursor into previousCursor successfully.
func TestEncodeValidBeforeCursor(t *testing.T) {
	// Given
	validBase64 := base64.StdEncoding.EncodeToString([]byte("test-cursor"))
	cursor := paginator.Cursor{Before: &validBase64}

	// When
	previousCursor, _, err := EncodeURLValues(cursor)

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, previousCursor)
	decoded, _ := base64.URLEncoding.DecodeString(previousCursor)
	assert.Equal(t, "test-cursor", string(decoded))
}

// Encode valid 'After' cursor into nextCursor successfully.
func TestEncodeValidAfterCursor(t *testing.T) {
	// Given
	validBase64 := base64.StdEncoding.EncodeToString([]byte("next-cursor"))
	cursor := paginator.Cursor{After: &validBase64}

	// When
	_, nextCursor, err := EncodeURLValues(cursor)

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, nextCursor)
	decoded, _ := base64.URLEncoding.DecodeString(nextCursor)
	assert.Equal(t, "next-cursor", string(decoded))
}

// Return empty strings for previousCursor and nextCursor when both cursors are nil.
func TestEncodeNilCursors(t *testing.T) {
	// Given
	cursor := paginator.Cursor{}

	// When
	previousCursor, nextCursor, err := EncodeURLValues(cursor)

	// Then
	assert.NoError(t, err)
	assert.Empty(t, previousCursor)
	assert.Empty(t, nextCursor)
}

// Handle nil 'Before' cursor and valid 'After' cursor correctly.
func TestHandleNilBeforeAndValidAfterCursor(t *testing.T) {
	// Given
	validBase64 := base64.StdEncoding.EncodeToString([]byte("next-cursor"))
	cursor := paginator.Cursor{After: &validBase64}

	// When
	previousCursor, nextCursor, err := EncodeURLValues(cursor)

	// Then
	assert.NoError(t, err)
	assert.Empty(t, previousCursor)
	assert.NotEmpty(t, nextCursor)
}

// Handle valid 'Before' cursor and nil 'After' cursor correctly.
func TestHandleValidBeforeAndNilAfterCursor(t *testing.T) {
	// Given
	validBase64 := base64.StdEncoding.EncodeToString([]byte("test-cursor"))
	cursor := paginator.Cursor{Before: &validBase64}

	// When
	previousCursor, nextCursor, err := EncodeURLValues(cursor)

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, previousCursor)
	assert.Empty(t, nextCursor)
}
