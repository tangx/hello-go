package main

import (
	"fmt"
	"testing"
	"time"
)

var retryChannel = make(map[string]chan string)

// var dnspodChan = make(chan string)
// var alidnsChan = make(chan string)

func Test_Main(t *testing.T) {

	go fn1("dnspod")
	go fn1("alidns")

	time.Sleep(2 * time.Second)

	for i := 0; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		retryChannel["dnspod"] <- "rktl.xyz"
	}

}

func fn1(name string) {
	ch := make(chan string)
	retryChannel[name] = ch

	for {
		time.Sleep(1 * time.Second)

		domain := <-retryChannel[name]
		fmt.Printf("channel [%s] -> %s\n", name, domain)
	}
}
