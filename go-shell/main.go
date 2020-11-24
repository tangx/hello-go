package main

import (
	"fmt"

	"github.com/progrium/go-shell"
)

func main() {
	// fmt.Println()
	// p := shell.Cmd("echo", "foobar").Pipe("wc", "-c").Pipe("awk", "'{print $1}'").Run()

	copy := shell.Cmd("cp").ErrFn()

	err := copy("/tmp/12123", "/tmp/abc")
	fmt.Println(err)
}
