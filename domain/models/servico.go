package models

type Servico struct {
	ItemListaServico       int32
	Discriminacao          string
	CodigoMunicipio        string
	ValorServicos          float32
	ValorDeducoes          float32
	ValorPis               float32
	ValorCofins            float32
	ValorInss              float32
	ValorIr                float32
	ValorCsll              float32
	IssRetido              float32
	ValorIss               float32
	ValorIssRetido         float32
	OutrasRetencoes        float32
	BaseCalculo            float32
	Aliquota               float32
	ValorLiquidoNfse       float32
	DescontoIncondicionado float32
	DescontoCondicionado   float32
}
