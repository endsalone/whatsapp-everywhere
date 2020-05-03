package whatsapp

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	log "github.com/sirupsen/logrus"
	"time"
)

func Connect() {
	wac, err := whatsapp.NewConn(1 * time.Second)
	if err != nil {
		log.Errorf("Error: %s", err)
	}

	wac.SetClientVersion(0, 6, 0)

	wac.AddHandler(&Handler{wac, uint64(time.Now().Unix())})

	session, err := login(wac)
	if err != nil {
		log.Errorf("Error when create a session: %v", err)
	}

	log.Printf("%s\n", sess)
}

func login(wac *whatsapp.Conn) (whatsapp.Session, error) {
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()

	session, err := wac.Login(qrChan)

	return session, err
}
