package funcdemo

import "testing"

func Test_FuncDemo(t *testing.T) {
	myfunc("k", "v")
}

func myfunc(k, v string) (bucket map[string]string) {
	// bucket = make(map[string]string) // 需要赋值

	bucket[k] = v

	return

	// === RUN   Test_FuncDemo
	// --- FAIL: Test_FuncDemo (0.00s)
	// panic: assignment to entry in nil map [recovered]
	// 	panic: assignment to entry in nil map
}
