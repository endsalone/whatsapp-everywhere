package rasa

import (
	"fmt"
	"github.com/endsalone/whatsapp-everywhere/config"
)

func conversationUrl() string {
	return fmt.Sprintf("%s/conversations", config.Bot.Url)
}
