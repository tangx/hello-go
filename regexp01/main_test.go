package regexp01

import (
	"fmt"
	"regexp"
	"testing"
)

// 正则表达式

func Test_Regexp(t *testing.T) {
	chars := `('|")`
	str := `"123'abc'456"`
	re := regexp.MustCompile(chars)
	s := re.ReplaceAllString(str, `\$1`)
	fmt.Println(s) // \"123\'abc\'456\"

}
