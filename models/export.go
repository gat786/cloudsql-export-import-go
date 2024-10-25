package models

import ()

type ExportContext struct {
	FileType         string   `json:"fileType"`
	URI              string   `json:"uri"`
	Databases        []string `json:"databases"`
	Offload          bool     `json:"offload"`
	SqlExportOptions struct {
		Clean    bool `json:"clean"`
		IfExists bool `json:"ifExists"`
	} `json:"sqlExportOptions"`
}

type ExportRequestBody struct {
	ExportContext ExportContext `json:"exportContext"`
}
