package bodyprocess

import (
	"encoding/json"
	"fmt"
	"github.com/endsalone/whatsapp-everywhere/internal/log"
	"io"
	"io/ioutil"
	"net/http"
)

func Process(body io.ReadCloser, converted interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Errorf("Error on read body: %s", err)
		return err
	}

	if len(data) == 0 {
		return nil
	}

	if err = json.Unmarshal(data, converted); err != nil {
		return err
	}

	return nil
}

func ProcessBodyError(res *http.Response) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("process body: %s", err)
	}

	return fmt.Errorf("status response %d with body '%s'", res.StatusCode, body)
}
