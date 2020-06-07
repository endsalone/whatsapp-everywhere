package rasa

import (
	"github.com/endsalone/whatsapp-everywhere/internal/bodyprocess"
	"github.com/endsalone/whatsapp-everywhere/internal/httpclient"
	"github.com/endsalone/whatsapp-everywhere/internal/log"
)

func SendConversation(message string, sender string) (*ConversationResponse, error) {
	url := conversationUrl()
	data := createRasaRequest(message, sender)

	body := bodyprocess.JsonParser{Data: data}
	response, err := httpclient.Post(url, body, nil, true)
	if err != nil {
		return nil, err
	}

	var conversationResponse ConversationResponse
	if err = bodyprocess.Process(response.Body, &conversationResponse); err != nil {
		log.Errorf("Error on process login response: %s", err)
		return nil, err
	}

	return &conversationResponse, nil
}
