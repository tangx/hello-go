package castconvert

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

type Student struct {
	Name string
	Age  int
}

func TestMain(t *testing.T) {
	zz := Student{Name: "zz"}
	zz.Age = cast.ToInt("123")

	fmt.Println(zz)
}
