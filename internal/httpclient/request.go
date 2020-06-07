package httpclient

import (
	"crypto/tls"
	"github.com/endsalone/whatsapp-everywhere/config"
	"github.com/endsalone/whatsapp-everywhere/internal/benchmark"
	"github.com/endsalone/whatsapp-everywhere/internal/bodyprocess"
	"net/http"
	"time"
)

var Client *http.Client
var InsecureClient *http.Client

func Setup() {
	Client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * time.Duration(10),
	}

	insecureTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	InsecureClient = &http.Client{
		Transport:     insecureTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * time.Duration(config.Server.HttpClientTimeout),
	}
}

func Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer benchmark.Request(benchmark.WithRequest(req))

	req.Header.Add("Accept", "application/json")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return Client.Do(req)
}

func Post(url string, body bodyprocess.Parser, headers map[string]string, isInsecure bool) (*http.Response, error) {
	reader, err := body.Parse()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	defer benchmark.Request(benchmark.WithRequest(req))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", body.ContentType())

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if isInsecure {
		return InsecureClient.Do(req)
	}

	return Client.Do(req)
}
