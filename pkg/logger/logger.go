package logger

import (
	"context"
	"strings"
	"time"

	"github.com/Kortivex/connected_roots/pkg/logger/commons"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represent the commons instance that you need to use in your implementation.
type Logger struct {
	logger     *zap.Logger
	level      string
	lvl        int
	tag        commons.Tag
	cid        commons.Cid
	ctx        *commons.Context
	middleware commons.MiddlewareFunc
}

// NewLogger
//
// Example:
//
//	loggerInstance := NewLogger("debug")
//	loggerInstance.WithTag(TagPlatformHttpserver)
//	loggerInstance.Info("success")
func NewLogger(lvl string) Logger {
	lvlInt := GetLevel(lvl)

	var logInstance *zap.Logger
	if lvl == "silent" {
		logInstance = zap.NewNop()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.MessageKey = "msg"
		cfg.EncoderConfig.LevelKey = "lvl"
		cfg.EncoderConfig.CallerKey = ""
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.Level = zap.NewAtomicLevelAt(zapcore.Level(lvlInt))
		cfg.Sampling = nil

		logInstance, _ = cfg.Build()
	}

	return Logger{logger: logInstance, level: lvl, lvl: lvlInt}
}

// New this method return a copy of the commons instance
//
// The main purpose of this Logger is instance a new commons with the same configuration but independent of the original.
func (log *Logger) New() Logger {
	newLogger := NewLogger(log.level)
	if log.tag != "" {
		newLogger.WithTag(log.tag)
	}
	if log.cid != "" {
		newLogger.WithCid(log.cid)
	}
	if log.ctx != nil {
		newLogger.WithCtx(*log.ctx)
	}
	if log.middleware != nil {
		newLogger.WithMiddleware(log.middleware)
	}
	return newLogger
}

func (log *Logger) NewEmpty() Logger {
	newLogger := NewLogger(log.level)
	return newLogger
}

func (log *Logger) GetLevel() (string, int) {
	return log.level, log.lvl
}

func (log *Logger) tsh() zap.Field {
	return zap.String("tsh", time.Now().UTC().Format("2006-01-02T15:04:05-0700"))
}

func (log *Logger) WithMiddleware(middleware commons.MiddlewareFunc) *Logger {
	log.middleware = middleware
	return log
}

func (log *Logger) WithTag(tag commons.Tag) *Logger {
	log.tag = tag
	tagField := zap.String("tag", string(tag))
	fieldsCustom := []zap.Field{tagField}
	log.logger = log.logger.With(fieldsCustom...)
	return log
}

func (log *Logger) WithCid(cid commons.Cid) *Logger {
	log.cid = cid
	cidField := zap.Any("cid", cid)
	fieldsCustom := []zap.Field{cidField}
	log.logger = log.logger.With(fieldsCustom...)
	return log
}

func (log *Logger) WithCtx(ctx commons.Context) *Logger {
	log.ctx = &ctx
	ctxField := zap.Any("ctx", ctx)
	fieldsCustom := []zap.Field{ctxField}
	log.logger = log.logger.With(fieldsCustom...)
	return log
}

func (log *Logger) Debug(msg string) {
	log.write("debug", msg)
}

func (log *Logger) DebugC(ctx context.Context, msg string) {
	log.write("debug", msg)
}

func (log *Logger) Info(msg string) {
	log.write("info", msg)
}

func (log *Logger) InfoC(ctx context.Context, msg string) {
	log.write("info", msg)
}

func (log *Logger) Warn(msg string) {
	log.write("warn", msg)
}

func (log *Logger) WarnC(ctx context.Context, msg string) {
	log.write("warn", msg)
}

func (log *Logger) Error(msg string) {
	log.write("error", msg)
}

func (log *Logger) ErrorC(ctx context.Context, msg string) {
	log.write("error", msg)
}

func (log *Logger) Fatal(msg string) {
	log.write("fatal", msg)
}

func (log *Logger) FatalC(ctx context.Context, msg string) {
	log.write("fatal", msg)
}

func (log *Logger) Panic(msg string) {
	log.write("panic", msg)
}

func (log *Logger) PanicC(ctx context.Context, msg string) {
	log.write("panic", msg)
}

func (log *Logger) write(lvl string, msg string) {
	if log.middleware != nil {
		var err error
		if msg, log.cid, log.tag, *log.ctx, err = log.middleware(msg, log.cid, log.tag, *log.ctx); err != nil {
			return
		}
	}

	switch lvl {
	case "debug":
		log.logger.Debug(msg, log.tsh())
	case "info":
		log.logger.Info(msg, log.tsh())
	case "warn":
		log.logger.Warn(msg, log.tsh())
	case "error":
		log.logger.Error(msg, log.tsh())
	case "fatal":
		log.logger.Fatal(msg, log.tsh())
	case "panic":
		log.logger.Panic(msg, log.tsh())
	}
}

func (log *Logger) Close() error {
	return log.logger.Sync()
}

func GetLevel(level string) int {
	switch strings.ToLower(level) {
	case "debug":
		return -1
	case "info":
		return 0
	case "warn":
		return 1
	case "error":
		return 2
	case "silent":
		return -2
	}
	return 0
}
