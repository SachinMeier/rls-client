package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	cli "github.com/urfave/cli"
)

func loadTLS(ctx *cli.Context) *http.Client {
	if ctx.GlobalIsSet(flagTLSPath) {
		tlsPath := ctx.GlobalString(flagTLSPath)
		return loadCertAndKey(tlsPath)
	} else {
		tlsPath := os.Getenv(rlsTLSPathKey)
		if tlsPath != "" {
			return loadCertAndKey(tlsPath)
		}
		return http.DefaultClient
	}
}

func loadCertAndKey(tlsPath string) *http.Client {
	clientCert := tlsPath + ".cert"
	clientKey := tlsPath + ".key"

	cert, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatalf("failed to load x509 certs: %v", err)
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{cert},
			},
		},
	}
}
