package string_test

import (
	"fmt"
	"strconv"
	"testing"
)

//转换int
func ExampleParseInt() {
	fmt.Println(strconv.ParseInt("123", 10, 64))
	fmt.Println(strconv.ParseUint("123", 10, 64))
	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.ParseInt("128", 10, 8))

	// Output:
	// 123 <nil>
	// 123 <nil>
	// 123 <nil>
	// 127 strconv.ParseInt: parsing "128": value out of range
}

//转换bool
func ExampleParseBool() {
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.FormatBool(true))
	fmt.Printf("%s\n", strconv.AppendBool([]byte("hello "), true))

	// Output:
	// true <nil>
	// true
	// hello true
}

//转换浮点数
func ExampleParseFloat() {
	fmt.Println(strconv.ParseFloat("12.3", 64))
	fmt.Println(strconv.FormatFloat(12.3, 'e', 3, 64))

	// Output:
	// 12.3 <nil>
	// 1.230e+01
}

//Quote
func ExampleQuote() {
	fmt.Println("This is", strconv.Quote("studygolang.com"), "website")
	fmt.Println(`This is "studygolang.com" website`)

	// Output:
	// This is "studygolang.com" website
	// This is "studygolang.com" website
}

//比较Itoah和sprintf的效率
func BenchmarkItoa(b *testing.B) {
	//比较fmt转换和Itoa转换效率
	for i := 0; i < b.N; i++ {
		strconv.Itoa(i) //49.8 ns/op
	}
}

//比较Itoah和sprintf的效率
func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i) //132 ns/op
	}
}