package uploadFiles

import "modTest/module"

type FilesRes struct {
	module.Resp
	Body string `json:"body"`
}

type FileImageSql struct {
	module.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Data string `json:"data"`
}

type FileImageDelete struct {
	module.Resp
	Body string `json:"body"`
}

type FileTest struct {
	Code string `json:"code"`
}
