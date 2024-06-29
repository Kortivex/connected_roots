package ulid

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

// Generate generate an ULID identifier.
func Generate() (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		return "", fmt.Errorf("%s: %w", "error generating ID", err)
	}

	return id.String(), nil
}
