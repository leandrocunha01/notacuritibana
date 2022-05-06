package models

import (
	"fmt"
	"testing"
	"time"
)

func TestModelNotaFiscal(t *testing.T) {
	prestador := Prestador{Cnpj: "34553142000106", InscricaoMunicipal: "010508640391"}
	tomador := Tomador{Cnpj: "48479236000189", RazaoSocial: "JMM Serviços de Logística"}
	servico := Servico{
		ItemListaServico:       1,
		Discriminacao:          "SERVIÇOS DE CONSULTORIA EM TI",
		CodigoMunicipio:        "4106902",
		ValorServicos:          7300.00,
		DescontoCondicionado:   0,
		ValorDeducoes:          0,
		ValorPis:               0,
		ValorCofins:            0,
		ValorInss:              0,
		ValorIr:                0,
		ValorCsll:              0,
		IssRetido:              0,
		ValorIss:               0,
		ValorIssRetido:         0,
		OutrasRetencoes:        0,
		BaseCalculo:            0.02,
		Aliquota:               0,
		ValorLiquidoNfse:       0,
		DescontoIncondicionado: 0.0,
	}
	endereco := Endereco{
		Endereco:        "Rua Santa Catarina",
		Numero:          "250",
		Bairro:          "Água Verde",
		Cep:             "80620120",
		Uf:              "PR",
		Complemento:     "Edifício Terra, AP 101",
		CodigoMunicipio: "4125506",
	}
	notaFiscal := NotaFiscal{
		Prestador:                &prestador,
		Tomador:                  &tomador,
		DataEmissao:              time.Now(),
		NaturezaOperacao:         1,
		RegimeEspecialTributacao: 1,
		OptanteSimplesNacional:   1,
		IncentivadorCultural:     0,
		Status:                   1,
		Servico:                  &servico,
		Endereco:                 &endereco,
	}

	fmt.Println(notaFiscal)
}
