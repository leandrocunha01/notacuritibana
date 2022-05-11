package util

import (
	"bytes"
	"text/template"
)

func TemplateXmlNfse(templateName string, model any) *bytes.Buffer {
	var fileBuffer bytes.Buffer
	paths := []string{
		Env("BASE_DIR") + "ConsultarNfse.xml",
		Env("BASE_DIR") + "ConsultaLoteRPS.xml",
		Env("BASE_DIR") + "NotaFiscal.xml",
		Env("BASE_DIR") + "ConsultarSituacaoLoteRPS.xml",
		Env("BASE_DIR") + "CancelarLoteNfe.xml",
	}
	t := template.Must(template.New(templateName).ParseFiles(paths...))

	err := t.Execute(&fileBuffer, model)
	if err != nil {
		panic(err)
	}
	return &fileBuffer
}
