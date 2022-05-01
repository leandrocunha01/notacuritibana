package main

import "C"
import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func s() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Renegotiation:      tls.RenegotiateOnceAsClient,
			InsecureSkipVerify: true},
	}
	//var file, _ = os.ReadFile("../files/fullchain.pem")
	var consultarXml, _ = os.ReadFile("../files/ConsultarNfse.xml")

	client := http.Client{Timeout: time.Duration(30) * time.Second, Transport: tr}
	req, err := http.NewRequest("POST", "https://piloto-iss.curitiba.pr.gov.br/nfse_ws/nfsews.asmx?WSDL", strings.NewReader(string(consultarXml)))
	req.Header = http.Header{
		"Host":         []string{"isscuritiba.curitiba.pr.gov.br"},
		"Content-Type": []string{"text/xml"},
		"SOAPAction":   []string{"http://www.e-governeapps2.com.br/ConsultarNfse"},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
