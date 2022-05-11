package models

import (
	"bytes"
	"github.com/leandrocunha01/notacuritibana/application"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
	"time"
)

type NotaFiscal struct {
	NumeroLote               int16
	QuantidadeRps            int
	Numero                   int16
	Serie                    int16
	Tipo                     int16
	DataEmissao              time.Time
	NaturezaOperacao         int
	RegimeEspecialTributacao int
	OptanteSimplesNacional   int
	IncentivadorCultural     int
	Status                   int
	Servico                  *Servico
	Prestador                *Prestador
	Tomador                  *Tomador
}

func (notaFiscal *NotaFiscal) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("NotaFiscal.xml", &notaFiscal)

	return application.NfseClient(xmlTemplate)
}

func (notaFiscal *NotaFiscal) Xml() *bytes.Buffer {
	xmlTemplate := util.TemplateXmlNfse("NotaFiscal.xml", &notaFiscal)
	return xmlTemplate
}
