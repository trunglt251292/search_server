package config

import (
	"github.com/Selly-Modules/mongodb"
	"go.elastic.co/apm/module/apmmongo"
)

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

// GetConnectOptions ...
func (dbCfg MongoCfg) GetConnectOptions() mongodb.Config {
	return mongodb.Config{
		Host:   dbCfg.Host,
		DBName: dbCfg.DBName,
		Standalone: &mongodb.ConnectStandaloneOpts{
			AuthMechanism: dbCfg.Mechanism,
			AuthSource:    dbCfg.Source,
			Username:      dbCfg.Username,
			Password:      dbCfg.Password,
		},
		Monitor: apmmongo.CommandMonitor(),
	}
}
