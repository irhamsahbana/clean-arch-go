package logger

import (
	"ca-boilerplate/bootstrap"

	"github.com/sirupsen/logrus"
)

func Log(fields logrus.Fields) *logrus.Logger {
	l := bootstrap.App.Log

	return l.WithFields(fields).Logger
}
