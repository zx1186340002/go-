package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	time string = "2006-01-02 15:04:05"
)

func initlogconfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time})
	}
}
func init() {
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()
	logrusLevel := viper.GetString("LOG_LEVEL")
	initlogconfig(logrusLevel)
}
