package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	if os.Getenv("IS_SINGLE_TEST") != "1" {
		return
	}
	os.Args = []string{"cmd", "server"}
	main()
}
