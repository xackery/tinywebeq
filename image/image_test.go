package image

import (
	"os"
	"testing"
)

func TestItemPreview(t *testing.T) {
	data, err := GenerateItemPreview(nil)
	if err != nil {
		t.Fatalf("GenerateItemPreview: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("no data")
	}
	err = os.WriteFile("test.png", data, 0644)
	if err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

}
