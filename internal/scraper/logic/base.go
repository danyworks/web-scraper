package logic

type HTMLElement struct {
	Index       int              `json:"id"`
	TagName     string           `json:"tag_name"`
	Description string           `json:"name"`
	VisibleText string           `json:"visible_text"`
	Properties  []HTMLProperties `json:"properties"`
}

type HTMLProperties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ScannerResults struct {
	Data []HTMLElement `json:"data"`
}
