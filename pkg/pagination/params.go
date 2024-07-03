package pagination

import (
	"encoding/base64"
	"net/http"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

const (
	defaultLimit    = 50
	defaultMaxLimit = 1000
	defaultOrder    = paginator.ASC
)

func getPaginatorParams(qParam *PaginatorParams) error {
	nextCursorDecoded, err := base64.URLEncoding.DecodeString(qParam.NextCursor)
	if err != nil {
		return ErrorPagination{Status: http.StatusBadRequest, Message: ErrMsgInvalidCursor}
	}
	qParam.NextCursor = base64.StdEncoding.EncodeToString(nextCursorDecoded)

	previousCursorDecoded, err := base64.URLEncoding.DecodeString(qParam.PreviousCursor)
	if err != nil {
		return ErrorPagination{Status: http.StatusBadRequest, Message: ErrMsgInvalidCursor}
	}
	qParam.PreviousCursor = base64.StdEncoding.EncodeToString(previousCursorDecoded)

	if qParam.Limit <= 0 {
		qParam.Limit = defaultLimit
	}

	if qParam.Limit > defaultMaxLimit {
		qParam.Limit = defaultMaxLimit
	}

	switch qParam.Order {
	case paginator.ASC:
		qParam.Order = paginator.ASC
	case paginator.DESC:
		qParam.Order = paginator.DESC
	default:
		qParam.Order = defaultOrder
	}

	return nil
}
