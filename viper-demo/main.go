package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/tangx/hello-go/viper-demo/internal"
)

func main() {
	// viper.SetDefault("ContentDir", "content2")
	// viper.SetDefault("LayoutDir", "layouts")

	// content := viper.Get("ContentDir")
	// fmt.Println(content)
	fmt.Println(viper.Get("config.filedir"))
	viper.GetInt("123")
}

func init() {
	// internal.Init()

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err.Error())
	}

}
