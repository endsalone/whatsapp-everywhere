package bodyprocess

import (
	"io"
	"net/url"
	"strings"
)

type UrlEncodedParser map[string]string

func (p UrlEncodedParser) Parse() (io.Reader, error) {
	data := url.Values{}

	for key, value := range p {
		data.Set(key, value)
	}

	return strings.NewReader(data.Encode()), nil
}

func (p UrlEncodedParser) ContentType() string {
	return "application/x-www-form-urlencoded"
}
