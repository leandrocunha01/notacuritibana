package models

import (
	"github.com/leandrocunha01/notacuritibana/appliccation"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
)

type ConsultarNfse struct {
	Prestador            *Prestador
	DataInicial          string
	DataFinal            string
	NumeroNfse           int32
	Tomador              *Tomador
	IntermediarioServico *IntermediarioServico
}

func (consultaNfse *ConsultarNfse) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("ConsultarNfse.xml", &consultaNfse)
	return appliccation.NfseClient(xmlTemplate)
}
