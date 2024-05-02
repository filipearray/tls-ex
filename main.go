package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

const (
	urlPath = "https://distopia-a1e2.savi2w.workers.dev/"
	customUserAgent = "arrayy"
)

func main() {
	tlsCustomConfig := tls.Config{
		MaxVersion: tls.VersionTLS12,
		MinVersion: tls.VersionTLS12,
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				con, err := tls.Dial(network, addr, &tlsCustomConfig)
				return con, err
			},
		},
	}

	request, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("User-Agent", customUserAgent)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
}