package commons

import (
	"errors"
	"strings"
)

type MiddlewareFunc func(message string, cid Cid, tag Tag, ctx Context) (string, Cid, Tag, Context, error)

func SkipExcludeUris(excludeUris []string) MiddlewareFunc {
	return func(message string, cid Cid, tag Tag, ctx Context) (string, Cid, Tag, Context, error) {
		for _, s := range excludeUris {
			if strings.HasPrefix(ctx.Request.Uri, s) {
				return message, cid, tag, ctx, errors.New(ctx.Request.Uri + " should not logged")
			}
		}

		return message, cid, tag, ctx, nil
	}
}
