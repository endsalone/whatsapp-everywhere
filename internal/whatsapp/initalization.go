package whatsapp

import (
	"encoding/gob"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/endsalone/whatsapp-everywhere/config"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Connect() {
	wac, err := whatsapp.NewConn(1 * time.Second)
	if err != nil {
		log.Errorf("Error: %s", err)
	}

	wac.SetClientVersion(0, 4, 2081)

	wac.AddHandler(&handler{
		Wac:       wac,
		starttime: uint64(time.Now().Unix()),
	})

	session, err := login(wac)
	if err != nil {
		log.Errorf("Error when create a session: %v", err)
	}

	pong, err := wac.AdminTest()
	if !pong || err != nil {
		log.Errorf("error pinging in: %v\n", err)
	}

	defer Connect()

	log.Info(session)
	if session.ServerToken != "" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		log.Warn("Safe shutdown")

		_, err := wac.Disconnect()
		if err != nil {
			log.Errorf("error disconnecting: %v", err)
		}
		log.Warn("Session finished")
	}
}

func login(wac *whatsapp.Conn) (whatsapp.Session, error) {
	session, err := readSession(wac)
	if err != nil {
		session, err := createQr(wac)
		if err != nil {
			return session, err
		}

	}

	return session, nil
}

func readSession(wac *whatsapp.Conn) (whatsapp.Session, error) {
	session := whatsapp.Session{}

	file, err := os.Open(config.Whatsapp.FileSession)
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}

	session, err = wac.RestoreWithSession(session)
	if err != nil {
		log.Errorf("Error when try to restore the session: %s", err)
	}
	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create(config.Whatsapp.FileSession)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}

	return nil
}

func createQr(wac *whatsapp.Conn) (whatsapp.Session, error) {
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
	}()

	session, err := wac.Login(qrChan)
	if err != nil {
		log.Errorf("Error when login: %s", err)
		return session, err
	}

	err = writeSession(session)
	if err != nil {
		log.Errorf("Error when write a new session: %s", err)
	}

	return session, err
}

func deleteSessionFile() {
	err := os.Remove(config.Whatsapp.FileSession)
	if err != nil {
		log.Errorf("Error when delete file")
	}
}
