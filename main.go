package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
)

func main() {

	var users = make(map[string]Person)

	p1 := Person{"zhang", 10}
	users["zhang"] = p1

	logger()
}

func logger() {

	logrus.Trace("logrus Trance = ", int(logrus.TraceLevel))
	logrus.Debug("logrus debug = ", int(logrus.DebugLevel))
	logrus.Info("logrus info = ", int(logrus.InfoLevel))
	logrus.Warn("logrus Warn = ", int(logrus.WarnLevel))
	fmt.Println("hello go")
	logrus.Error("logrus error = ", int(logrus.ErrorLevel))
	logrus.Fatal("logrus Fatal = ", int(logrus.FatalLevel))
	logrus.Panic("logrus Panic = ", int(logrus.PanicLevel))
}
func Greet(who string) {
	logrus.Infof("Hello, %s", who)

	p := Person{"zhangsan", 20}
	query.Values(p)
}

type Person struct {
	Name string
	Age  int
}
