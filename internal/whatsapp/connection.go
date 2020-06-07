package whatsapp

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
)

type handler struct {
	Wac       *whatsapp.Conn
	Contacts  []whatsapp.Contact
	Chat      []whatsapp.Chat
	starttime uint64
}

var Handler handler

type contact struct {
	Name  string
	Phone string
}

func (h *handler) HandleError(err error) {
	//if e, ok := err.(*whatsapp.ErrConnectionFailed); ok {
	//	log.Printf("Connection failed, underlying error: %v", e.Err)
	//	log.Println("Waiting 30sec...")
	//	<-time.After(30 * time.Second)
	//	log.Println("Reconnecting...")
	//	err := h.Wac.Restore()
	//	if err != nil {
	//		log.Fatalf("Restore failed: %v", err)
	//	}
	//} else {
	//	log.Errorf("Handler Error: %s\n", err)
	//}
}

func (h *handler) ShouldCallSynchronously() bool {
	return true
}

func (h *handler) HandleContactList(contactList []whatsapp.Contact) {
	Handler.Contacts = contactList
}

func (h *handler) HandleChatList(contacts []whatsapp.Chat) {
	Handler.Chat = contacts
	fmt.Println("Chat", len(Handler.Chat))
	for i, contact := range Handler.Contacts {
		fmt.Println("Contacts on Chat: ", contact, i)
	}
}
