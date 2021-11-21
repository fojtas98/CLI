package template

import (
	"io/ioutil"
	"os"
	"text/template"
)

func CreateFromTamplate() *template.Template {
	currentWorkingDirectory, _ := os.Getwd()
	file, err := os.Open(currentWorkingDirectory + "/template/template.txt")
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
