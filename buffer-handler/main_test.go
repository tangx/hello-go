package bufferhandler

import (
	"fmt"
	"os"
	"testing"

	"github.com/fatih/color"
)

func Test_WriteTo(t *testing.T) {

	// WriteTo("git-template")
	f, err := os.OpenFile("git-template.sh", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		color.Red(fmt.Sprint(err))
	}
	defer f.Close()
	WriteTo(os.Stdout, "gt")
	WriteTo(f, "gt")
}

func Test_bufferWriteTo(t *testing.T) {
	bufferWriteTo()
}
