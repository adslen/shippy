package log

import (
	"os"

	"github.com/op/go-logging"
)

var (
	defaultLogger *logging.Logger
)

var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02T15:04:05.000} %{shortfile} %{shortfunc} ? %{level:.4s} %{id:03x} %{color:reset} %{message}`,
)

func init() {
	defaultLogger = logging.MustGetLogger("shippy")

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	//backendFormatter := logging.GlogFormatter

	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backendLeveled, backendFormatter)

}

func DefaultLogger() *logging.Logger {
	if defaultLogger != nil {
		return defaultLogger
	}

	defaultLogger = logging.MustGetLogger("shippy")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.DEBUG, "")
	logging.SetBackend(backendLeveled, backendFormatter)

	return defaultLogger
}

type Logger interface {
}

func L() *logging.Logger {
	return defaultLogger
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	DefaultLogger().Fatal(args...)
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	DefaultLogger().Fatalf(format, args...)
}

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
func Panic(args ...interface{}) {
	DefaultLogger().Panic(args...)

}

// Panicf is equivalent to l.Critical followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	DefaultLogger().Panicf(format, args...)
}

// Critical logs a message using CRITICAL as log level.
func Critical(args ...interface{}) {
	DefaultLogger().Critical(args...)
}

// Criticalf logs a message using CRITICAL as log level.
func Criticalf(format string, args ...interface{}) {
	DefaultLogger().Criticalf(format, args...)
}

// Error logs a message using ERROR as log level.
func Error(args ...interface{}) {
	DefaultLogger().Error(args...)
}

// Errorf logs a message using ERROR as log level.
func Errorf(format string, args ...interface{}) {
	DefaultLogger().Errorf(format, args...)
}

// Warning logs a message using WARNING as log level.
func Warning(args ...interface{}) {
	DefaultLogger().Warning(args...)
}

// Warningf logs a message using WARNING as log level.
func Warningf(format string, args ...interface{}) {
	DefaultLogger().Warningf(format, args...)
}

// Notice logs a message using NOTICE as log level.
func Notice(args ...interface{}) {
	DefaultLogger().Notice(args...)
}

// Noticef logs a message using NOTICE as log level.
func Noticef(format string, args ...interface{}) {
	DefaultLogger().Noticef(format, args...)
}

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	DefaultLogger().Info(args...)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	DefaultLogger().Infof(format, args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	DefaultLogger().Debug(args...)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	DefaultLogger().Debugf(format, args...)
}
