package usecase

import (
	"bytes"
	"github.com/leandrocunha01/notacuritibana/application"
	"github.com/leandrocunha01/notacuritibana/domain/models"
	"github.com/leandrocunha01/notacuritibana/util"
	"net/http"
)

type CancelarLoteNfe struct {
	Prestador       *models.Prestador
	NumeroNfse      int32
	CodigoMunicipio int32
}

func (cancelarLoteNfe *CancelarLoteNfe) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("CancelarLoteNfe.xml", &cancelarLoteNfe)

	return application.NfseClient(xmlTemplate)
}

func (cancelarLoteNfe *CancelarLoteNfe) Xml() *bytes.Buffer {
	xmlTemplate := util.TemplateXmlNfse("CancelarLoteNfe.xml", &cancelarLoteNfe)
	return xmlTemplate
}
