package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

var gs = []byte(`timeout: 3s`)

type Game struct {
	Timeout time.Duration `json:"timeout"`
}

func main() {
	var g Game

	_ = yaml.Unmarshal(gs, &g)

	spew.Dump(g)
	fmt.Println(time.Now().String())
	time.Sleep(g.Timeout)
	fmt.Println(time.Now().String())
}
