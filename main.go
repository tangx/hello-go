package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/tangx/go-querystring/query"
	"github.com/tangx/hello-go/elastic"
)

func main() {

	fmt.Println(color.BlueString("hello"))
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
