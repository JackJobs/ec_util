package string_test

import (
	"bytes"
	"fmt"
)

func ExampleBytes() {
	fmt.Println(bytes.Contains([]byte("hello"), []byte("el")))
	fmt.Println(bytes.Fields([]byte("hello world")))
	// Output:
	// true
	// [[104 101 108 108 111] [119 111 114 108 100]]
}
