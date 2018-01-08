package map_test

import (
	"fmt"
)

func ExampleArray() {
	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// Output:
	// [1 2 3]
	// [1 2 3 0 0 0 0 0 0 0]
	// [1 2 3]
}

func ExampleSlice() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	fmt.Println(s1)
	fmt.Println(s2)
	// Output:
	// [1 2 3]
	// [0 0 0]
}

func ExampleMap() {
	m1 := make(map[string]int)
	m1["one"] = 1
	fmt.Println(m1)
	delete(m1, "one")
	// Output:
	// map[one:1]
}


