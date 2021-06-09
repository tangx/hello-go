package forlabel

import (
	"testing"
)

func Test_Main(t *testing.T) {

	i := 0

OutLoop:
	for {
		i++
		if i%2 == 0 {
			continue OutLoop // continue 只能用于循环
		}

		if i%3 == 0 {
			break OutLoop
		}
	}

	switch i {
	case 5:
		break OutLoop
	}
}
