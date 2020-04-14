package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Trace("logrus Trance = ", int(logrus.TraceLevel))
	logrus.Debug("logrus debug = ", int(logrus.DebugLevel))
	logrus.Info("logrus info = ", int(logrus.InfoLevel))
	logrus.Warn("logrus Warn = ", int(logrus.WarnLevel))
	fmt.Println("hello go")
	logrus.Error("logrus error = ", int(logrus.ErrorLevel))
	logrus.Fatal("logrus Fatal = ", int(logrus.FatalLevel))
	logrus.Panic("logrus Panic = ", int(logrus.PanicLevel))
}
