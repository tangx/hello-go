package deferdemo

import (
	"fmt"
	"os"
	"testing"
)

func Test_DeferDemo(t *testing.T) {
	deferRun()
}

func deferRun() {
	defer fmt.Println("defer: hello")

	// panic("some error") // ok defer
	// log.Fatal("fatal exit") // no defer
	os.Exit(1) // no defer

}
