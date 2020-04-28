package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/hello-go/elastic"
)

func main() {

	// logger.PrintLogLevel()
	elastic.ESMain()

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
