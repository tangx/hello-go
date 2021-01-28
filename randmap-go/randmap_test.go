package randmapper

import (
	"fmt"
	"testing"
)

var m = make(map[string]bool)

func Test_Randomap(t *testing.T) {

	m["http://127.0.0.1:8080"] = true
	m["socks://127.0.0.1:8080"] = true
	m["https://baidu.com:443"] = true

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	return k
	// }

	for i := 0; i < 1; i++ {
		fmt.Println(randmap(m))
	}

}

func randmap(m map[string]bool) string {
	for k := range m {
		return k
	}

	return ""
}
