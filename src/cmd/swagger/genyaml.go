package main

import (
	"os"
	"text/template"

	"github.com/cowk8s/harbor/src/lib/config/metadata"
)

const cfgTemplate = `  Configurations:
    type: object
    properties: {{ range .Items }}
      {{ .Name }}:
        type: {{ .Type }}
        description: {{ .Description }} 
        x-omitempty: true
        x-isnullable: true{{ end }}
`

const responseTemplate = `  ConfigurationResponse:
    type: object
    properties: {{ range .Items }}
      {{ .Name }}:
        $ref: '#/definitions/{{ .Type }}'
        description: {{ .Description }} {{ end }}
`

type document struct {
	Items []templateItem
}

type templateItem struct {
	Name        string
	Type        string
	Description string
}

func userCfgItems(isResponse bool) []templateItem {
	items := make([]templateItem, 0)
	for _, i := range metadata.ConfigList {
		item := templateItem{
			Name: i.Name,
		}
		items = append(items, item)
	}
	return items
}

type yamlFile struct {
	Name       string
	IsResponse bool
	TempName   string
}

func main() {
	l := []yamlFile{
		{"configurations.yml", false, cfgTemplate},
		{"configurationsResponse.yml", true, responseTemplate},
	}
	for _, file := range l {
		f, err := os.Create(file.Name)
		if err != nil {
			panic(err)
		}
		doc := document{
			Items: userCfgItems(file.IsResponse),
		}
		tmpl, err := template.New("test").Parse(file.TempName)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(f, doc)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}
