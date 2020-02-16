package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/snwfdhmp/errlog"
)

func CheckErr(e error) {
	if e != nil {
		errlog.Debug(e)
		panic(e)
	}
}

// Log logs if debug env var is set at true
func Log(ctxt context.Context, level string, msg ...interface{}) {
	switch level {
	case "info":
		logrus.Infoln(msg...)
	case "warn":
		logrus.Warnln(msg...)
	case "error":
		logrus.Errorln(msg...)
	default:
		logrus.Infoln(msg...)
	}
}
