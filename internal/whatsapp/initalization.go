package whatsapp

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"log"
	"time"
)

func Conection() {
	wac, err := whatsapp.NewConn(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code %s\n", <-qrChan)
	}()

	sess, sessErr := wac.Login(qrChan)
	if sessErr != nil {
		log.Fatal(sessErr)
	}

	log.Printf("%s\n", sess)
}
