package notacuritibana

import (
	"bytes"
	"github.com/leandrocunha01/notacuritibana/application"
	"github.com/leandrocunha01/notacuritibana/models"
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
	Servico                  *models.Servico
	Prestador                *models.Prestador
	Tomador                  *models.Tomador
}

func (notaFiscal *NotaFiscal) Send() *http.Response {
	xmlTemplate := util.TemplateXmlNfse("NotaFiscal.xml", &notaFiscal)

	return application.NfseClient(xmlTemplate)
}

func (notaFiscal *NotaFiscal) Xml() *bytes.Buffer {
	xmlTemplate := util.TemplateXmlNfse("NotaFiscal.xml", &notaFiscal)
	return xmlTemplate
}

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
