package global

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jkunii/crossJ/helper"
	"gopkg.in/caarlos0/env.v2"
)

type Config struct {
	ServerDomain               string `env:"SERVER_DOMAIN" envDefault:"http://localhost:1323"`
	LogLevel                   int    `env:"LOG_LEVEL" envDefault:"5"`
	Port                       string `env:"PORT" envDefault:"1323"`
	ShowBanner                 bool   `env:"SHOW_BANNER" envDefault:"false"`
	LogLevelType               string `env:"LOG_LEVEL_TYPE" envDefault:"debug"`
	LogFormater                string `env:"LOG_FORMATER" envDefault:"txt"`
	UserName                   string `env:"BAISC_USER" envDefault:"123"`
	Secret                     string `env:"BASIC_SECRET" envDefault:"123"`
	MongoAnalyticsActive       bool   `env:"MONGO_ANALYTICS_ACTIVE" envDefault:"true"`
	MongoHosts                 string `env:"MONGO_HOSTS" envDefault:"localhost,localhost"`
	MongoUser                  string `env:"MONGO_USER" envDefault:"rosie_usr"`
	MongoUserSecret            string `env:"MONGO_USER_SECRET" envDefault:"rosie_usr"`
	MongoDatabase              string `env:"MONGO_DATABASE" envDefault:"rosie"`
	MongoCertificateCollection string `env:"MONGO_CERTIFICATECOLLECTION" envDefault:"certificate"`
}

var (
	Cfg    Config
	logger = logrus.New()
)

func (r Config) Init() {

	err := env.Parse(&Cfg)
	if err != nil {
		fmt.Println("Error getting environment variables")
		fmt.Println(err)
	}

	//Logger
	if Cfg.LogFormater == "json" {
		logger.Formatter = new(logrus.JSONFormatter)
	} else {
		logger.Formatter = new(logrus.TextFormatter) // default
	}

	logger.Level, err = logrus.ParseLevel(Cfg.LogLevelType)
	helper.PanicErr(err)

}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func LogLevel() logrus.Level {
	return logger.Level
}
