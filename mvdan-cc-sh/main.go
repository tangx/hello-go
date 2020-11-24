package main

import (
	"fmt"

	"mvdan.cc/sh/v3/shell"
)

func main() {
	env := func(name string) string {
		switch name {
		case "HOME":
			return "/home/user"
		}
		return "" // leave the rest unset
	}
	out, _ := shell.Expand("No place like $HOME", env)
	fmt.Println(out)

	out, _ = shell.Expand("Some vars are ${missing:-awesome}", env)
	fmt.Println(out)

	out, _ = shell.Expand("Math is fun! $((12 * 34))", nil)
	fmt.Println(out)
	// Output:
	// No place like /home/user
	// Some vars are awesome
	// Math is fun! 408

	cmd := `cp -a /tmp/1230sdf0uasdf /tmp/23sf0sdf0`
	out, err := shell.Expand(cmd, nil)
	fmt.Println(out)
	if err != nil {
		panic(err)
	}
}
