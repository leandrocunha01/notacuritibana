package usecase

import (
	"github.com/leandrocunha01/notacuritibana/application"
	"github.com/leandrocunha01/notacuritibana/domain/models"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
)

type ConsultarLoteRps struct {
	Prestador *models.Prestador
	Protocolo string
}

type ConsultarSituacaoLoteRps struct {
	Prestador *models.Prestador
	Protocolo string
}

func (consultarLoteRps *ConsultarLoteRps) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("ConsultaLoteRPS.xml", &consultarLoteRps)

	return application.NfseClient(xmlTemplate)
}

func (ConsultarSituacaoLoteRps *ConsultarSituacaoLoteRps) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("ConsultarSituacaoLoteRPS.xml", &ConsultarSituacaoLoteRps)

	return application.NfseClient(xmlTemplate)
}
