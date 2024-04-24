package image

import (
	"os"
	"testing"

	"github.com/xackery/tinywebeq/model"
	"gopkg.in/yaml.v3"
)

func TestItemPreview(t *testing.T) {

	item := &model.Item{
		Name:    "Singing Short Sword",
		Classes: 1,
		Races:   1,
	}

	r, err := os.Open("20542.yaml")
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("os.Open: %v", err)
	}
	if err == nil {
		defer r.Close()
		err = yaml.NewDecoder(r).Decode(item)
		if err != nil {
			t.Fatalf("yaml.NewDecoder.Decode: %v", err)
		}

	}

	data, err := GenerateItemPreview(item)
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
