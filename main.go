package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/hello-go/logger"
	"github.com/tangx/hello-go/web"
)

func main() {

	web.NewServer()

	logger.PrintLogLevel()

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
