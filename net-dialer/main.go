package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := "172.16.255.156:5432"
	n := 2
	interval := 2
	conn, err := net.DialTimeout("tcp", addr, time.Duration(n)*time.Second)
	if err != nil {
		fmt.Printf("failed to connect to tcp://%s, retry after %d seconds :%v",
			addr, interval, err)
		time.Sleep(time.Duration(interval) * time.Second)
	}
	if err = conn.Close(); err != nil {
		fmt.Printf("failed to close the connection: %v", err)
	}
	fmt.Printf("succes dial tcp://%s \n", addr)
}
