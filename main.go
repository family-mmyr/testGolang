package main

import (
	"github.com/family-mmyr/testGolang/logger"
	"github.com/sirupsen/logrus"
)

func init() {
	f := logger.TextFormatter{}
	f.RequestID = "123456789"
	logrus.SetFormatter(&f)
}

func main() {
	logrus.WithFields(logrus.Fields{"id": "aaa"}).Error("invalid id")
}
