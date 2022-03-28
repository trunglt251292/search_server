package config

type Config struct {
	Port          string `env:"PORT"`
	Elasticsearch struct {
		URL      string `env:"ES_URL,required"`
		Username string `env:"ES_USER"`
		Password string `env:"ES_PWD"`
	}
	Mongo  MongoCfg   `env:",prefix=MONGO_"`
	Nats   NatsConfig `env:",prefix=NATS_"`
	ApiKey string     `env:"ES_SERVICE_APIKEY"`
}

// NatsConfig ...
type NatsConfig struct {
	URL      string `env:"URL,required"`
	Username string `env:"USER"`
	Password string `env:"PWD"`
}

// MongoCfg ...
type MongoCfg struct {
	Host      string `env:"URI"`
	DBName    string `env:"NAME"`
	Mechanism string `env:"MECHANISM"`
	Source    string `env:"SOURCE"`
	Username  string `env:"USERNAME"`
	Password  string `env:"PWD"`
}
