package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}
