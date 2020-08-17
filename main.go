package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	p := os.Getenv("PHASE")

	if p == "DEBUG" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		// デフォルトはテキスト形式のため、JSONフォーマットに切替える
		logrus.SetFormatter(&logrus.JSONFormatter{})

		// Warn以上ののログを出力する
		logrus.SetLevel(logrus.InfoLevel)
	}

}

func main() {
	logrus.WithFields(logrus.Fields{
		"log_version": "v1",
		"request_id":  "aaa",
	}).Info("operation")

	logrus.WithFields(logrus.Fields{
		"log_version": "v1",
		"request_id":  "aaa",
	}).Debug("debug")
}
