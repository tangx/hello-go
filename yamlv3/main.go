package main

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

var (
	sourceYaml = `
# user address city
USER__ADDRESS__CITY: sichuan
# user address number
USER__ADDRESS__NUMBER: 1
# user Gender
USER__GENDER: true

# time clock
TIME__CLOCK: 3m

`
)

func main() {
	convert()
	// userDefine()
	// transfer()
}

func convert() {
	t := yaml.Node{}

	err := yaml.Unmarshal([]byte(sourceYaml), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// spew.Dump(t)

	for _, v := range t.Content {
		spew.Dump(v)
		// fmt.Println(v.Content[0].HeadComment)
		// v.Content[0].LineComment = "adfasdfaskdjfalsjkd"
		v.Content[0].FootComment = "12j12l3l"
	}

	b, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func userDefine() {
	t := yaml.Node{
		Kind: 4,
	}

	// tk := yaml.Node{}

	key := &yaml.Node{
		Kind:        8,
		Value:       "USER__ADDRESS__CITY",
		HeadComment: "# USER__ADDRESS__CITY: 用户居住地址",
	}

	value := &yaml.Node{
		Kind:  8,
		Value: "chengdu",
	}

	t.Content = append(t.Content, key, value)

	b, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}

func transfer() {
	user := struct {
		Name   string
		Age    int32
		Gender bool
		Time   time.Duration
	}{"zhangsan", 20, true, 30 * time.Hour}

	b, _ := yaml.Marshal(&user)
	fmt.Println(string(b))

}
