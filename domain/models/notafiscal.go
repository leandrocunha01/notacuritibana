package models

import (
	"github.com/leandrocunha01/notacuritibana/appliccation"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
	"time"
)

type NotaFiscal struct {
	DataEmissao              time.Time
	NaturezaOperacao         int
	RegimeEspecialTributacao int
	OptanteSimplesNacional   int
	IncentivadorCultural     int
	Status                   int
	Servico                  *Servico
	Prestador                *Prestador
	Tomador                  *Tomador
	Endereco                 *Endereco
}

func (notaFiscal *NotaFiscal) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("NotaFiscal.xml", &notaFiscal)

	return appliccation.NfseClient(xmlTemplate)
}
