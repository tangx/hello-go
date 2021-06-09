package main

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type User struct {
	Name      string
	Age       int
	Addresses *Addr
}

type Addr struct {
	Street string
	Number int
}

const (
	filename = "demo.gob"
)

func encoding() {

	user := User{
		Name: "zhangsan",
		Age:  20,
		Addresses: &Addr{
			Street: "chansha",
			Number: 1000,
		},
	}

	File, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(&user); err != nil {
		fmt.Println()
	}
}

func main() {
	// encoding()
	decoding()
}

func decoding() {
	var user User
	File, _ := os.Open(filename)
	D := gob.NewDecoder(File)
	err := D.Decode(&user)
	if err != nil {
		logrus.Error(err)
	}
	// fmt.Println(user.Name)
	addr := user.Addresses
	fmt.Println(addr.Street)
}
