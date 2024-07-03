package pagination

import (
	"encoding/base64"
	"net/http"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

func EncodeURLValues(cursor paginator.Cursor) (string, string, error) {
	var previousCursor, nextCursor string
	if cursor.Before != nil {
		beforeDecoded, err := base64.StdEncoding.DecodeString(*cursor.Before)
		if err != nil {
			return "", "", ErrorPagination{Status: http.StatusInternalServerError, Message: ErrMsgInternalServer}
		}
		previousCursor = base64.URLEncoding.EncodeToString(beforeDecoded)
	}

	if cursor.After != nil {
		afterDecoded, err := base64.StdEncoding.DecodeString(*cursor.After)
		if err != nil {
			return "", "", ErrorPagination{Status: http.StatusInternalServerError, Message: ErrMsgInternalServer}
		}
		nextCursor = base64.URLEncoding.EncodeToString(afterDecoded)
	}

	return previousCursor, nextCursor, nil
}
