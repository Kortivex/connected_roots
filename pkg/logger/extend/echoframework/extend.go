package echoframework

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"
	"github.com/labstack/echo/v4"
)

// ErrorHandlerMiddlewareLogger is a compatible error handler middleware for `httpserver` library.
func ErrorHandlerMiddlewareLogger(loggerInstance *logger.Logger) func(h echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(h echo.HTTPErrorHandler) echo.HTTPErrorHandler {
		return func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}

			errorLogger := loggerInstance.NewEmpty()
			errorLogger.WithTag(commons.TagPlatformHttpserver)
			cid := commons.GetCIDFromEcho(c)

			eS, ok := err.(commons.ErrorI)
			if !ok {
				eS = commons.NewDefaultErrorS(err)
			}
			errorContext := commons.NewContext().FromError(eS)

			errorLogger.WithCtx(*errorContext).WithCid(cid).Error(eS.Error())

			h(err, c)
		}
	}
}

// PreMiddlewareLogger is a compatible middleware for echo.
//
// This middleware generate a new logger adding `cid` from echo.HeaderXCorrelationID or echo.HeaderXRequestID.
//
// Important: This middleware should be set with Pre method:
//
//	echoInstance.Pre(PreMiddlewareLogger(logger, true))
//
// If you set withRequestContext flag to `true` the request is adding to log trace.
//
// In the echo handler you can use this logger getting from echo:
//
//	echoInstance.GET("/ok", func(c echo.Context) error {
//	    c.Logger().Info("ok from common echo.context")
//	    return c.String(http.StatusOK, "ok")
//	})
func PreMiddlewareLogger(loggerInstance *logger.Logger, withRequestContext bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			loggerPerRequest := loggerInstance.New()
			loggerPerRequest.WithCid(commons.GetCIDFromEcho(c))

			if withRequestContext {
				c.Set("withRequestContext", true)

				//loggerPerRequest.WithCtx(generateRequestCtx(logCtx, c.Request(), c.Response(), c.RealIP(), 0))
			}

			c.Set("logger", loggerPerRequest)
			// c.SetLogger(NewLogger(&loggerPerRequest))

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}

// PostMiddlewareLogger is a compatible middleware for echo.
//
// # This middleware
//
// Important: This middleware should be set with Use method:
//
//	echoInstance.Use(PostMiddlewareLogger(slowThreshold, "warn"))
//
// slowThreshold param will use for emit log trace is the duration of request raise this duration.
//
// slowThresholdLevel param define the log level for emit the log trace. ("silent","error","warn","info","debug").
func PostMiddlewareLogger(slowThreshold time.Duration, slowThresholdLevel string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			loggerInstance := c.Get("logger").(logger.Logger)
			logCtx := *commons.NewContext()

			err := next(c)
			if err != nil {
				c.Error(err)
				var errGeneric commons.ErrorI
				var ok bool
				if errGeneric, ok = err.(*commons.ErrorS); ok {
					logCtx.FromError(errGeneric)
				} else {
					response := map[string]any{}
					response["size"] = c.Response().Size
					response["status"] = c.Response().Status
					response["header"] = c.Response().Header()
					errGeneric = commons.NewErrorS(c.Response().Status, "", response, err)
					logCtx.FromError(errGeneric)
				}
			}

			stop := time.Now()

			logCtx = generateRequestCtx(logCtx, c.Request(), c.Response(), c.RealIP(), stop.Sub(start))
			loggerInstance.WithCtx(logCtx)

			switch {
			case c.Response().Status >= http.StatusInternalServerError:
				loggerInstance.Error("server_error")
			case c.Response().Status >= http.StatusBadRequest:
				loggerInstance.Warn("client_error")
			case c.Response().Status >= http.StatusMultipleChoices:
				loggerInstance.Info("redirection")
			default:
				if err != nil {
					loggerInstance.Error("error")
				} else {
					if slowThresholdLevel != "" && slowThresholdLevel != "silent" && stop.Sub(start) > slowThreshold {
						switch slowThresholdLevel {
						case "error":
							loggerInstance.Error("success but slow")
						case "warn":
							loggerInstance.Warn("success but slow")
						case "info":
							loggerInstance.Info("success but slow")
						case "debug":
							loggerInstance.Debug("success but slow")
						}
					} else {
						// c.Logger().Debug("success", zap.Any("ctx", logCtx))
					}
				}
			}

			return nil
		}
	}
}

func generateRequestCtx(logCtx commons.Context, req *http.Request, res *echo.Response, ip string, latency time.Duration) commons.Context {
	latencyString := strconv.FormatInt(int64(latency), 10)
	latencyHuman := latency.String()
	bytesIn := 0
	cl := req.Header.Get(echo.HeaderContentLength)
	if cl != "" {
		bytesIn, _ = strconv.Atoi(cl)
	}

	logCtx.AddRequestContext(commons.RequestContext{
		Host:      req.Host,
		Method:    req.Method,
		Uri:       req.RequestURI,
		RemoteIp:  ip,
		UserAgent: req.UserAgent(),
		Latency:   latencyString,
		LatencyH:  latencyHuman,
		Status:    res.Status,
		BytesIn:   int64(bytesIn),
		BytesOut:  res.Size,
	})

	return logCtx
}
