package util

import (
	"bytes"
	"text/template"
)

func TemplateXmlNfse(templateName string, model any) *bytes.Buffer {
	var fileBuffer bytes.Buffer
	paths := []string{
		"/home/leandro/GolandProjects/notacuritibana/templates/ConsultarNfse.xml",
		"/home/leandro/GolandProjects/notacuritibana/templates/ConsultaLoteRPS.xml",
	}

	t := template.Must(template.New(templateName).ParseFiles(paths...))

	err := t.Execute(&fileBuffer, model)
	if err != nil {
		panic(err)
	}
	return &fileBuffer
}
