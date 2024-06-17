package gorm

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

// Logger implement gorm.Logger interface
//
// You can use NewLogger for instanced Logger struct.
type Logger struct {
	iLogger               *logger.Logger
	slowThreshold         time.Duration
	skipErrRecordNotFound bool
}

// NewLogger instanced Logger struct
//
//	 mainLogger := NewLogger("info")
//	 gormLogger := mainLogger.New()
//	 logger := gormLogger.WithTag(commons.TagPlatformGorm)
//	 pgConfig := postgres.Config{
//	     ...
//			Logger: &NewLogger(logger, time.Millisecond, false),
//	     ...
//		}
func NewLogger(logger *logger.Logger, slowThreshold time.Duration, skipErrRecordNotFound bool) Logger {
	return Logger{
		iLogger:               logger,
		slowThreshold:         slowThreshold,
		skipErrRecordNotFound: skipErrRecordNotFound,
	}
}

func (l *Logger) LogMode(gormLog.LogLevel) gormLog.Interface {
	return l
}

func (l *Logger) Info(ctx context.Context, s string, args ...interface{}) {
	log := l.iLogger.New()

	cid := ctx.Value("cid")
	if cid != nil {
		log.WithCid(commons.Cid(cid.(string)))
	}

	message := fmt.Sprintf(s, args)
	logCtx := commons.NewContext()
	logCtx.AddDump(args)

	log.WithCtx(*logCtx)

	log.Info(message)
}

func (l *Logger) Warn(ctx context.Context, s string, args ...interface{}) {
	log := l.iLogger.New()

	cid := ctx.Value("cid")
	if cid != nil {
		log.WithCid(commons.Cid(cid.(string)))
	}

	message := fmt.Sprintf(s, args)
	logCtx := commons.NewContext()
	logCtx.AddDump(args)

	log.WithCtx(*logCtx)

	log.Warn(message)
}

func (l *Logger) Error(ctx context.Context, s string, args ...interface{}) {
	log := l.iLogger.New()

	cid := ctx.Value("cid")
	if cid != nil {
		log.WithCid(commons.Cid(cid.(string)))
	}

	message := fmt.Sprintf(s, args)
	logCtx := commons.NewContext()
	e := commons.NewErrorS(0, message, nil, nil)
	logCtx.FromError(e)

	log.WithCtx(*logCtx)

	log.Error("gorm error")
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	loggerInstance := l.iLogger.New()

	stackLines := strings.Split(string(debug.Stack()), "\n")
	caller := strings.Replace(stackLines[len(stackLines)-2], "\t", "", 1)

	logCtx := commons.NewContext()
	logCtx.AddQueryContext(commons.QueryContext{
		Caller:  caller,
		Query:   sql,
		Rows:    rows,
		Errors:  nil,
		Latency: float64(elapsed.Nanoseconds()) / 1e6,
	})

	cid := ctx.Value("cid")
	if cid != nil {
		loggerInstance.WithCid(commons.Cid(cid.(string)))
	}

	loggerInstance.WithCtx(*logCtx)

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.skipErrRecordNotFound) {
		e := commons.NewErrorS(0, err.Error(), nil, err)
		logCtx.FromError(e)
		loggerInstance.Error("gorm error")
		return
	}

	if l.slowThreshold != 0 && elapsed > l.slowThreshold {
		loggerInstance.Warn("success but slow")
		return
	}

	// loggerInstance.Debug("success")
}
