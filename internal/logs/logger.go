package logs

import (
	"context"
	"fmt"
	"golang-rate-limit/internal/constants"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	logger    *log.Logger
	logPrefix string
	tags      []string
}

// buildLog generates log's message
func (l *Logger) buildLog(ctx context.Context, message string) string {
	return fmt.Sprintf("[%s] %s %s", l.logPrefix, message, l.getTags(ctx))
}

// getTags find the request ID from the context
func (l *Logger) getTags(ctx context.Context) string {
	stringBuilder := strings.Builder{}
	for _, tag := range l.tags {
		value := ctx.Value(tag)

		tag := fmt.Sprintf("[%s:%s]", tag, value)

		stringBuilder.WriteString(tag)
	}

	return stringBuilder.String()
}

// WithLogger This method is only for testing purposes and should never be used
func (l *Logger) WithLogger(logger *log.Logger) {
	l.logger = logger
}

func New(logPrefix string) Logger {
	l := log.New()

	l.SetFormatter(formatter())
	l.SetOutput(setStdout())
	l.SetLevel(setLevel())

	return Logger{
		logger:    l,
		logPrefix: logPrefix,
		tags: []string{
			constants.RequestID,
		},
	}
}

// formatter return the logrus formatter set in the environment variable
// Valid values:
// - json
// - text
// Default: logrus.JSONFormatter
func formatter() log.Formatter {
	envVal := os.Getenv(constants.EnvLogFormatter)
	switch strings.ToLower(envVal) {
	case "text":
		return new(log.TextFormatter)
	case "json":
		return new(log.JSONFormatter)
	default:
		return new(log.JSONFormatter)
	}
}

// setStdout return where to store logs set in the environment variable
// Those options are:
// - stdout
// - stderr
// Default: os.Stdout
func setStdout() io.Writer {
	envVal := os.Getenv(constants.EnvLogStdOut)
	switch strings.ToLower(envVal) {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		return os.Stdout
	}
}

// setLevel return the logrus level configured by environment variable
// Valid values:
// - panic
// - fatal
// - error
// - warn
// - warning
// - info
// - debug
// - trace
func setLevel() log.Level {
	envVal := os.Getenv(constants.EnvLogLevel)

	level, err := log.ParseLevel(envVal)
	if err != nil {
		return log.DebugLevel
	}

	return level
}

func (l *Logger) Error(ctx context.Context, message string, err error) {
	log := l.buildLog(ctx, l.errorMessage(message, err))

	l.logger.Error(log)
}

func (l *Logger) ErrorWithoutContext(message string, err error) {
	log := l.errorMessage(message, err)

	l.logger.Error(log)
}

func (l *Logger) errorMessage(message string, err error) string {
	return fmt.Sprintf("%s - Error: %s", message, err)
}

func (l *Logger) Info(ctx context.Context, message string) {
	l.logger.Info(l.buildLog(ctx, message))
}

func (l *Logger) InfoWithoutContext(message string) {
	l.logger.Info(message)
}
