package site

import "github.com/xackery/tinywebeq/config"

type BaseData struct {
	Title       string
	BaseURL     string
	GoogleTag   string
	Description string
	ImageURL    string
}

func BaseDataInit(titleSuffix string) BaseData {
	return BaseData{
		Title:       "tinywebeq | " + titleSuffix,
		Description: titleSuffix,
		BaseURL:     config.Get().Site.BaseURL,
		GoogleTag:   config.Get().Site.GoogleTag,
	}
}

func BaseDataInitWithImage(titleSuffix, imageURL string) BaseData {
	return BaseData{
		Title:       "tinywebeq | " + titleSuffix,
		Description: titleSuffix,
		BaseURL:     config.Get().Site.BaseURL,
		GoogleTag:   config.Get().Site.GoogleTag,
		ImageURL:    imageURL,
	}
}
