package templates

import (
	"embed"
	"io/ioutil"
	"text/template"
)

//go:embed *.txt
var templateFs embed.FS

func CreateFromTamplate(templateName string) *template.Template {
	file, err := templateFs.Open(templateName + ".txt")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	getWebsite, err := template.New("template").Parse(string(b))
	if err != nil {
		panic(err)
	}
	return getWebsite
}
