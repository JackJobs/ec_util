package string_test

import (
	"fmt"
	"unicode"
)

func ExampleUnicode() {
	single := '\u0015'
	fmt.Println(unicode.IsControl(single)) //true
	single = '\uef35'
	fmt.Println(unicode.IsControl(single)) //false

	digit := rune('1')
	fmt.Println(unicode.IsDigit(digit))  //true
	fmt.Println(unicode.IsNumber(digit)) //true
	letter := rune('Ⅷ')
	fmt.Println(unicode.IsDigit(letter))  //false
	fmt.Println(unicode.IsNumber(letter)) //true

	fmt.Println(rune('张'))
	fmt.Println(len([]rune("Go编程")))

	// Output:
	// true
	// false
	// true
	// true
	// false
	// true
	// 24352
	// 4

}
