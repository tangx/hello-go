package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func doClientWork(client *rpc.Client, i string) {
	go func() {
		fmt.Println("in client work go func()")
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 130, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		color.Green("watch:", keyChanged)
	}()

	// time.Sleep(3 * time.Second)
	// fmt.Println("time to sleep")
	key := "tangxin" + i
	value := "tangxin-value" + i

	fmt.Printf("key = %s ; value = %s\n", key, value)

	err := client.Call(
		"KVStoreService.Set", [2]string{key, value},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}

func main() {

	client, err := rpc.Dial("tcp", "127.0.0.1:11234")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 3)

		doClientWork(client, strconv.Itoa(i))
	}

}
