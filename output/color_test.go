package output

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func Test_ColorOutput(t *testing.T) {
	s := color.BlueString("hello %s", "golang")
	fmt.Println(s)
}
