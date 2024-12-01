package logs

import (
	log "github.com/sirupsen/logrus"
)

func Debug(file map[string]interface{}, msg string) {
	log.WithFields(file).Debug(msg)
}
