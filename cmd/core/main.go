package main

import (
	"github.com/endsalone/whatsapp-everywhere/config"
	"github.com/endsalone/whatsapp-everywhere/internal/httpclient"
	"github.com/endsalone/whatsapp-everywhere/internal/whatsapp"
)

func main() {
	config.Setup()
	httpclient.Setup()

	whatsapp.Connect()
}
