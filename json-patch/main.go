package main

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

func main() {
	// Let's create a merge patch from these two documents...
	original := []byte(`{"name": "John", "age": 24, "height": 3.21}`)
	target := []byte(`{"name": "Jane", "age": 24}`)

	patch, err := jsonpatch.CreateMergePatch(original, target)
	if err != nil {
		panic(err)
	}

	// Now lets apply the patch against a different JSON document...

	alternative := []byte(`{"name": "Tina", "age": 28, "height": 3.75}`)
	modifiedAlternative, _ := jsonpatch.MergePatch(alternative, patch)

	fmt.Printf("patch document:   %s\n", patch)                      // {"height":null,"name":"Jane"}
	fmt.Printf("updated alternative doc: %s\n", modifiedAlternative) // {"age":28,"name":"Jane"}

	/* 说明:
	patch document:          {"height":null,"name":"Jane"}
	updated alternative doc: {"age":28,"name":"Jane"}
		1. patch 中不涉及 age, 因此 age=28 不变。
		2. patch 中 name=Jane, 因此 Tina -> Jane。
		3. patch height=nill, 因此 删除 height
	*/

}
