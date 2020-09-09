package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

type KVStoreServcie struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreServcie() *KVStoreServcie {
	return &KVStoreServcie{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreServcie) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreServcie) Set(kv [2]string, reply *struct{}) error {
	fmt.Println("Set be calld")
	defer fmt.Println("out set")
	p.mu.Lock()

	fmt.Println("Set get lock")
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]
	color.Yellow(key, value)

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}
	return nil
}

func (p *KVStoreServcie) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) //buffered

	// p.mu.Lock()  /* lock 会与 Set 争抢，造成阻塞*/
	p.filter[id] = func(key string) { ch <- key }
	// defer p.mu.Unlock()
	// spew.Dump(ch)

	spew.Dump(p.m)
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}

	// return nil
}

func main() {

	_ = rpc.RegisterName("KVStoreService", NewKVStoreServcie())

	listener, err := net.Listen("tcp", ":11234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("connect accept")
		if err != nil {
			log.Fatal("Accept error", err)
		}

		fmt.Println("new client")
		go rpc.ServeConn(conn)
	}
}
