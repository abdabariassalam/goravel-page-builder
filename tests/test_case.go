package tests

import (
	"github.com/goravel/framework/testing"

	"page_builder/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
