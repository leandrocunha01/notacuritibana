package models

import "time"

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
