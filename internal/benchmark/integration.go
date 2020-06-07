package benchmark

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func WithRequest(request *http.Request) (string, time.Time) {
	return fmt.Sprintf("%s: %s", request.Method, request.URL), time.Now()
}

func Request(s string, startTime time.Time) {
	took := time.Since(startTime)
	log.Infof("Time took to request [%s]: %s", s, took)
}
