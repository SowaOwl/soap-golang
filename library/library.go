package library

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/tiaguinho/gosoap"
	"net/http"
	"os"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	certPath := "new_cert.pem"

	caCert, err := os.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &Client{
		http: &http.Client{Transport: transport},
	}
}

func (c *Client) SendRequest(url string, methodName string, data map[string]interface{}) ([]byte, error) {
	client, err := gosoap.SoapClient(url, c.http)
	if err != nil {
		return nil, err
	}

	resp, err := client.Call(methodName, data)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
