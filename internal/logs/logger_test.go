package logs_test

import (
	"context"
	"errors"
	"golang-rate-limit/internal/logs"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func createLogger(prefix string) (*logs.Logger, *test.Hook) {
	mock, hook := test.NewNullLogger()

	l := logs.New(prefix)
	l.WithLogger(mock)

	return &l, hook
}

func createContext() context.Context {
	return context.WithValue(context.Background(), "request_id", "12341234")
}

func TestLogger_Error(t *testing.T) {
	logger, hook := createLogger("some-component")

	logger.Error(createContext(), "error in component", errors.New("error in api call"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "[some-component] error in component - Error: error in api call [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_ErrorWithoutContext(t *testing.T) {
	logger, hook := createLogger("some-component")

	logger.ErrorWithoutContext("error log", errors.New("error in api call"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "error log - Error: error in api call", hook.LastEntry().Message)
}

func TestLogger_Info(t *testing.T) {
	logger, hook := createLogger("some-component")

	logger.Info(createContext(), "info log")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "[some-component] info log [request_id:12341234]", hook.LastEntry().Message)
}

func TestLogger_InfoWithoutContext(t *testing.T) {
	logger, hook := createLogger("some-component")

	logger.InfoWithoutContext("info log")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "info log", hook.LastEntry().Message)
}
