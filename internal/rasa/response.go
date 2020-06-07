package rasa

type ConversationResponse struct {
	Data conversationDataResponse `json:"data"`
}

type conversationDataResponse struct {
	Type       string                         `json:"type"`
	Attributes conversationAttributesResponse `json:"attributes"`
}

type conversationAttributesResponse struct {
	Message string `json:"message"`
}
