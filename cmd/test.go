package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair("/home/leandro/GolandProjects/notacuritibana/files/cert.pem",
		"/home/leandro/GolandProjects/notacuritibana/files/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		Renegotiation:      tls.RenegotiateOnceAsClient,
	}
	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	var consultarXml, _ = os.Open("/home/leandro/GolandProjects/notacuritibana/files/ConsultarNfse.xml")
	fmt.Println(consultarXml)
	resp, err := client.Post("https://piloto-iss.curitiba.pr.gov.br/nfse_ws/nfsews.asmx?WSDL", "text/xml", consultarXml)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
