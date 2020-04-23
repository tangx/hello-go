package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/hello-go/logger"
)

func main() {

	var users = make(map[string]Person)

	p1 := Person{"zhang", 10}
	users["zhang"] = p1

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
