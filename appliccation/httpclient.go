package appliccation

import (
	"crypto/tls"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func NfseClient(xmlModel io.Reader) *http.Response {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured .env file. Err: %s", err)
	}
	certificate := os.Getenv("CERT")
	keyfile := os.Getenv("CERT_KEY")
	cert, err := tls.LoadX509KeyPair(certificate, keyfile)
	if err != nil {
		log.Fatal(err)
	}
	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true, // Ignora problemas de instalação da cadeia ICP
		Renegotiation:      tls.RenegotiateOnceAsClient,
	}
	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	resp, err := client.Post("https://piloto-iss.curitiba.pr.gov.br/nfse_ws/nfsews.asmx?WSDL", "text/xml", xmlModel)
	if err != nil {
		log.Fatalf("Client error. Err: %s", err)
	}
	return resp
}
