package output

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/tangx/colorlog"
)

func Test_ColorOutput(t *testing.T) {
	s := color.BlueString("hello %s", "golang")
	fmt.Println(s)

	color.Red("hello %s", "golang")
}

func Test_ColorlogOutput(t *testing.T) {
	colorlog.Info("hello, %s", "golang")
}
