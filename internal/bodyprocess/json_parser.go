package bodyprocess

import (
	"bytes"
	"encoding/json"
	"io"
)

type JsonParser struct {
	Data interface{}
}

func (p JsonParser) Parse() (io.Reader, error) {
	jsondata, err := json.Marshal(p.Data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsondata), nil
}

func (p JsonParser) ContentType() string {
	return "application/json"
}
