package site

type BaseData struct {
	Title string
}

func BaseDataInit(titleSuffix string) BaseData {
	return BaseData{
		Title: "tinywebeq | " + titleSuffix,
	}
}
