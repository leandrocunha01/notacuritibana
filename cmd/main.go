package main

import (
	"github.com/leandrocunha01/notacuritibana/domain/models"
	"io/ioutil"
	"log"
)

func main() {
	prestador := models.Prestador{Cnpj: "34553142000106", InscricaoMunicipal: "010508640391"}
	tomador := models.Tomador{Cpf: "00687334292"}
	consulta := models.ConsultarNfse{Tomador: &tomador, Prestador: &prestador}

	data, err := ioutil.ReadAll(consulta.Send().Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
