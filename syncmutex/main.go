package main

import "sync"

type Counter struct {
	Count1 uint64

	sync.Mutex
	Count2 uint64
	Count3 uint64
}
