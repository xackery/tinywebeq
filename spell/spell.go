package spell

import "fmt"

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
	return nil
}
