package usecase

import (
	"github.com/leandrocunha01/notacuritibana/appliccation"
	"github.com/leandrocunha01/notacuritibana/domain/models"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
)

type ConsultarLoteRps struct {
	Prestador *models.Prestador
	Protocolo string
}

func (consultarLoteRps *ConsultarLoteRps) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("ConsultaLoteRPS.xml", &consultarLoteRps)

	return appliccation.NfseClient(xmlTemplate)
}
