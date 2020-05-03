package whatsapp

import (
	"github.com/Rhymen/go-whatsapp"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Wac       *whatsapp.Conn
	starttime uint64
}

func (h Handler) HandleError(err error) {
	log.Errorf("Error when handler: %v\n", err)
}
