package config

type WhatsappConfig struct {
	FileSession string
}

var Whatsapp *WhatsappConfig

func WhatsappSetup() {
	Whatsapp = &WhatsappConfig{
		FileSession: getEnv("SESSION_FILE", "whatsapp-session") + ".gob",
	}
}
