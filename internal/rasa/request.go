package rasa

import (
	"io"
)

type conversationRequest struct {
	Data dataRequest `json:"data"`
}

func (c conversationRequest) Parse() (io.Reader, error) {
	panic("implement me")
}

func (c conversationRequest) ContentType() string {
	panic("implement me")
}

type dataRequest struct {
	Type       string            `json:"type"`
	Attributes attributesRequest `json:"attributes"`
}

type attributesRequest struct {
	Msg    string `json:"msg"`
	Sender string `json:"sender"`
}

func createRasaRequest(message string, sender string) *conversationRequest {
	return &conversationRequest{
		Data: dataRequest{
			Type: "conversations",
			Attributes: attributesRequest{
				Msg:    message,
				Sender: sender,
			},
		},
	}
}
