package main

import (
	"os"

	"github.com/alecthomas/chroma/quick"
)

const (
	Dockerfile = `FROM alpine:3.12
ADD go-bin /usr/local/bin/go-bin
RUN chmod +x /usr/local/bin/go-bin
ENTRYPOINT ["/usr/local/bin/go-bin"]
`
)

func main() {

	if err := quick.Highlight(
		os.Stdout, Dockerfile,
		"Docker", "terminal", "vim"); err != nil {
		panic(err)
	}

}
