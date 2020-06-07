package config

type BotConfig struct {
	Url string
}

var Bot *BotConfig

func BotSetup() {
	Bot = &BotConfig{
		Url: getEnv("BOT_URL", "http://localhost:3000"),
	}
}
