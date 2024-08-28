package item

import (
	"fmt"
)

var (
	isInitialied bool
)

func Init() error {
	if isInitialied {
		return nil
	}
	isInitialied = true
	err := viewInit()
	if err != nil {
		return fmt.Errorf("viewInit: %w", err)
	}
	err = searchInit()
	if err != nil {
		return fmt.Errorf("searchInit: %w", err)
	}
	return nil
}
