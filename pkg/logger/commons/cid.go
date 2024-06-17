package commons

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Cid string

func GetCIDFromEcho(c echo.Context) Cid {
	req := c.Request()
	res := c.Response()
	id := req.Header.Get(echo.HeaderXCorrelationID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXCorrelationID)
		if id == "" {
			id = req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
		}
	}

	cid := Cid(id)

	return cid
}

func GetCIDFromContext(c context.Context) Cid {
	requestID := c.Value("requestID").(string)
	if requestID == "" {
		requestID = c.Value(echo.HeaderXCorrelationID).(string)
		if requestID == "" {
			requestID = c.Value(echo.HeaderXRequestID).(string)
		}
	}

	return Cid(requestID)
}
