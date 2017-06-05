package main

import (
	"testing"
	"os"
)

// TODO: get rid of this global, need an easier way to pass data store to handlers
var service Service

func TestMain(m *testing.M) {
	service.InitService()

	// run all tests
	code := m.Run()
	os.Exit(code)
}