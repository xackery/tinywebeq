package quest

import (
	"fmt"

	"github.com/xackery/tinywebeq/quest/parse"
)

var (
	isInitialized bool
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true
	err := viewInit()
	if err != nil {
		return fmt.Errorf("viewInit: %w", err)
	}

	err = searchInit()
	if err != nil {
		return fmt.Errorf("searchInit: %w", err)
	}

	err = parse.Init()
	if err != nil {
		return fmt.Errorf("questParseInit: %w", err)
	}
	return nil
}
