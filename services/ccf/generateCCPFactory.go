package ccf

import (
	"fmt"
	"io"
	"text/template"
)

type Organization struct {
	OrgName string `json:"OrgName"`
	OrgNum  string `json:"OrgNum"`
	CAPort  string `json:"CAPort"`
	P0Port  string `json:"P0Port"`
	P1Port  string `json:"P1Port"`
}

type ConfigInfo struct {
	Organizations []Organization `json:"Organizations"`
}

//GenerateConfigTxTemplate To write the config to the w according to the template tpl
func GenerateCCPTemplate(configInfo ConfigInfo, tpl *template.Template, w io.Writer) error {
	err := tpl.Execute(w, configInfo)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
