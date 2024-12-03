package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	time string = "2006-01-02 15:04:05"
)

var JwtSignKey string

func initlogconfig(logLevel string) {
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time})
	}
}
func init() {
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("JWT_SIGN_KEY", "SMALL_ZHOU")
	viper.AutomaticEnv()
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	logrusLevel := viper.GetString("LOG_LEVEL")
	initlogconfig(logrusLevel)
}
