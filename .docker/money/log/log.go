package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
}

// TODO: implement log interface & builder pattern

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to message errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// Level type
type Level uint32

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	logger.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	logger.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	logger.SetLevel(logrus.Level(level))
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *logrus.Entry {
	return logger.WithFields(logrus.Fields(fields))
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info logs a message at level info on the standard logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Print logs a message at level print on the standard logger.
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Warn logs a message at level warn on the standard logger.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error logs a message at level error on the standard logger
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Panic logs a message at level panic on the standard logger and exit program.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Debugf logs a message at level Info on the standard logger.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Infoln ...
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Println ...
func Println(args ...interface{}) {
	logger.Println(args...)
}

// Warnln ...
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Errorln ...
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

// Panicln ...
func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}

// Fatalln ...
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

// Logger get logger instance
func Logger() *logrus.Logger {
	return logger
}
