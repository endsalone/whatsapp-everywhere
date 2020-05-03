package whatsapp

import (
	"encoding/gob"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	log "github.com/sirupsen/logrus"
	"os"
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

	log.Debugf("Session value: %s", session)

	defer Connect()

	<-time.After(1 * time.Second)
}

func login(wac *whatsapp.Conn) (whatsapp.Session, error) {
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()

	session, err := wac.Login(qrChan)

	return session, err
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}

	file, err := os.Open("this-is-my-session.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}

	return session, nil
}
