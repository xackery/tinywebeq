package item

import "fmt"

var (
	isInitialied bool
)

func Init() error {
	if isInitialied {
		return nil
	}
	isInitialied = true
	err := previewInit()
	if err != nil {
		return fmt.Errorf("previewInit: %w", err)
	}
	err = viewInit()
	if err != nil {
		return fmt.Errorf("viewInit: %w", err)
	}
	return nil
}
