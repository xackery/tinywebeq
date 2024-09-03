package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("IS_SINGLE_TEST") != "1" {
		return
	}
	os.Args = []string{"cmd", "server"}
	main()
}
