package main

import (
	"github.com/leandrocunha01/notacuritibana/domain/models"
	"github.com/leandrocunha01/notacuritibana/usecase"
	"io/ioutil"
	"log"
)

func main() {
	prestador := models.Prestador{Cnpj: "34553142000106", InscricaoMunicipal: "010508640391"}

	consulta := usecase.ConsultarLoteRps{Protocolo: "189899898898998", Prestador: &prestador}

	data, err := ioutil.ReadAll(consulta.Send().Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
