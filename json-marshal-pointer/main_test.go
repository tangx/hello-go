package jsonmarshalpointer

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"gopkg.in/yaml.v3"
)

// 结构体定义指针， json 进行 unmarshal

func Test_Main(t *testing.T) {
	p := person{}

	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", p)

	// fmt.Println(*p.Name)
	// fmt.Println(p.Name)
	// fmt.Println(p.Gender)

	b, err := yaml.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)

}

type person struct {
	Name   *string `yaml:"name,omitempty" json:"name,omitempty"`
	Age    *int    `yaml:"age,omitempty" json:"age,omitempty"`
	Gender *bool   `yaml:"gender,omitempty" json:"gender,omitempty"`
}

// var data = `{"name":"zhangsan","age":18,"gender":true}`
var data = `{"age":18,"gender":true}`
