package echoframework

import (
	"io"
	"sync/atomic"

	"github.com/Kortivex/connected_roots/pkg/logger"
	echoLog "github.com/labstack/gommon/log"
	"github.com/mattn/go-colorable"
)

// Logger implement echo.Logger interface
//
// You can use NewEchoLogger for instanced Logger struct.
type Logger struct {
	iLogger *logger.Logger
	output  io.Writer
	prefix  string
	level   uint32
}

// NewLogger instanced Logger struct
//
//	mainLogger := NewLogger("info")
//	echoLogger := mainLogger.New()
//	commons := echoLogger.WithTag(commons.TagPlatformEcho)
//	echoInstance := echo.New()
//	echoInstance.Logger = NewLogger(commons)
func NewLogger(logger *logger.Logger) *Logger {
	// iLogger := commons.New()
	_, lvl := logger.GetLevel()
	return &Logger{
		iLogger: logger, // iLogger.WithTag("echo_logger"),
		output:  output(),
		prefix:  "echoLogger",
		level:   uint32(lvl),
	}
}

// GetSWLogger get reference to commons instance.
func (l *Logger) GetSWLogger() *logger.Logger {
	return l.iLogger
}

func (l *Logger) Output() io.Writer {
	return l.output
}

func (l *Logger) SetOutput(w io.Writer) {
	l.output = w
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

func (l *Logger) Level() echoLog.Lvl {
	return echoLog.Lvl(atomic.LoadUint32(&l.level))
}

func (l *Logger) SetLevel(level echoLog.Lvl) {
	atomic.StoreUint32(&l.level, uint32(level))
}

func (l *Logger) SetHeader(h string) {
	// not apply
}

func (l *Logger) Print(i ...interface{}) {
	l.iLogger.Debug(i[0].(string))
}

func (l *Logger) Printf(_ string, arg ...interface{}) {
	l.iLogger.Debug(arg[0].(string))
}

func (l *Logger) Printj(j echoLog.JSON) {
	l.iLogger.Debug("")
}

func (l *Logger) Debug(i ...interface{}) {
	l.iLogger.Debug(i[0].(string))
}

func (l *Logger) Debugf(_ string, arg ...interface{}) {
	l.iLogger.Debug(arg[0].(string))
}

func (l *Logger) Debugj(j echoLog.JSON) {
	l.iLogger.Debug("")
}

func (l *Logger) Info(i ...interface{}) {
	l.iLogger.Info(i[0].(string))
}

func (l *Logger) Infof(_ string, arg ...interface{}) {
	l.iLogger.Info(arg[0].(string))
}

func (l *Logger) Infoj(j echoLog.JSON) {
	l.iLogger.Info("")
}

func (l *Logger) Warn(i ...interface{}) {
	l.iLogger.Warn(i[0].(string))
}

func (l *Logger) Warnf(_ string, arg ...interface{}) {
	l.iLogger.Warn(arg[0].(string))
}

func (l *Logger) Warnj(j echoLog.JSON) {
	l.iLogger.Warn("")
}

func (l *Logger) Error(i ...interface{}) {
	l.iLogger.Error(i[0].(string))
}

func (l *Logger) Errorf(_ string, arg ...interface{}) {
	l.iLogger.Error(arg[0].(string))
}

func (l *Logger) Errorj(j echoLog.JSON) {
	l.iLogger.Error("")
}

func (l *Logger) Fatal(i ...interface{}) {
	l.iLogger.Fatal(i[0].(string))
}

func (l *Logger) Fatalf(_ string, arg ...interface{}) {
	l.iLogger.Fatal(arg[0].(string))
}

func (l *Logger) Fatalj(j echoLog.JSON) {
	l.iLogger.Fatal("")
}

func (l *Logger) Panic(i ...interface{}) {
	l.iLogger.Panic(i[0].(string))
}

func (l *Logger) Panicf(_ string, arg ...interface{}) {
	l.iLogger.Panic(arg[0].(string))
}

func (l *Logger) Panicj(j echoLog.JSON) {
	l.iLogger.Panic("")
}

// output return a writer compatible with echo.Logger.
func output() io.Writer {
	return colorable.NewColorableStdout()
}
