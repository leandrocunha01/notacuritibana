package usecase

import (
	"github.com/leandrocunha01/notacuritibana/application"
	"github.com/leandrocunha01/notacuritibana/domain/models"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
)

type ConsultarNfse struct {
	Prestador            *models.Prestador
	DataInicial          string
	DataFinal            string
	NumeroNfse           int32
	Tomador              *models.Tomador
	IntermediarioServico *models.IntermediarioServico
}

func (consultaNfse *ConsultarNfse) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("ConsultarNfse.xml", &consultaNfse)

	return application.NfseClient(xmlTemplate)
}
