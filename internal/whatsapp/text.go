package whatsapp

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	"github.com/endsalone/whatsapp-everywhere/internal/rasa"
)

func (h *handler) HandleTextMessage(message whatsapp.TextMessage) {
	if message.Info.FromMe || h.starttime > message.Info.Timestamp {
		return
	}

	conversation, err := rasa.SendConversation(message.Text, message.Info.RemoteJid)
	if err != nil {
		fmt.Printf("Conversation Error: %s", err)
	}

	answer := conversation.Data.Attributes.Message

	fmt.Printf("%v", conversation)
	err = sendMessage(message, h.Wac, answer)
	if err != nil {
		fmt.Printf("ERROU: %s", err)
	}
	//for _, v := range h.Wac.Store.Contacts {
	//	fmt.Println("Nome:", v.Name, "Jid", v.Jid, "Short:", v.Short, "Notify", v.Notify)
	//}
	//log.Info(len(h.Wac.Store.Contacts))
}

func sendMessage(message whatsapp.TextMessage, wac *whatsapp.Conn, answer string) error {
	_, err := wac.Read(message.Info.RemoteJid, message.Info.Id)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: message.Info.RemoteJid,
		},
		ContextInfo: whatsapp.ContextInfo{
			QuotedMessageID: message.Info.Id,
			QuotedMessage: &proto.Message{
				Conversation: &message.Text,
			},
			Participant: message.Info.RemoteJid,
		},
		Text: answer,
	}
	_, err = wac.Send(msg)
	if err != nil && err.Error() != "sending message timed out" {
		fmt.Printf("%s", err)
		return err
	}

	return nil
}
