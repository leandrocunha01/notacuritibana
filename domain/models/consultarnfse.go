package models

import (
	"github.com/leandrocunha01/notacuritibana/appliccation"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
)

type ConsultarNfse struct {
	Empresa     *Empresa
	DataInicial string
	DataFinal   string
}

func (consultaNfse *ConsultarNfse) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("ConsultarNfse.xml", &consultaNfse)
	return appliccation.NfseClient(xmlTemplate)
}
