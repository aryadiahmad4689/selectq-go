package utils

import (
	"bytes"
	"html/template"
	"log"

	"github.com/aryadiahmad4689/selectq-go/src/entity"
)

func Parse(name, tplt string, t entity.DataQuery) string {
	tmpl, err := template.New(name).Parse(tplt)
	if err != nil {
		log.Fatal(err.Error())
	}
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, t); err != nil {
		log.Fatal(err.Error())
	}
	return buffer.String()
}
