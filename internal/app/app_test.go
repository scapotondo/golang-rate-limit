package app

import (
	"fmt"
	"net/http/httptest"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (suite *AppSuite) TestInjectDependencies() {
	app := NewApp()
	app.Setup()
	handler := app.Engine

	require.NotNil(suite.T(), handler)
}

func (suite *AppSuite) TestCreateServer() {
	app := NewApp()
	app.Setup()
	handler := app.Engine
	testServer := httptest.NewServer(handler)

	require.NotNil(suite.T(), testServer)
	testServer.Close()
}

func (suite *AppSuite) TestWaitForShutdown() {
	app := NewApp()
	server := app.createServer()

	go func() {
		fmt.Println("waiting 2 seconds")
		time.Sleep(2 * time.Second)
		fmt.Println("closing")
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	}()

	app.waitForShutdownSignal(server)
}

func (suite *AppSuite) TestInitServer() {
	app := NewApp()

	go func() {
		fmt.Println("waiting 2 seconds")
		time.Sleep(2 * time.Second)
		fmt.Println("closing")
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	}()

	app.InitServer()
}
