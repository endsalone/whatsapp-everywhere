package config

type ServerConfig struct {
	HttpClientTimeout int
}

var Server *ServerConfig

func ServerSetup() {
	Server = &ServerConfig{
		HttpClientTimeout: getIntEnv("SERVER_TIMEOUT", 10),
	}
}
