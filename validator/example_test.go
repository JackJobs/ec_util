package validator_test

import (
	"fmt"
	"regexp"
)

func ExampleCompile() {
	r, err := regexp.Compile(`hello`)
	if err != nil {
		fmt.Println("there is a problem with your regexp")
		return
	}
	if r.MatchString("hello world") == true {
		fmt.Println("match")
	} else {
		fmt.Println("not match")
	}

	// Output: match
}

func ExampleMustCompile() {
	//MustCompile
	if regexp.MustCompile(`hello`).MatchString("hello world") == true {
		fmt.Println("match")
	} else {
		fmt.Println("not match")
	}

	// Output: match
}

func ExamplePattern() {
	// \w ---- [A-Za-z0-9]
	r, _ := regexp.Compile(`H\wllo`)
	fmt.Printf("\\H: %v\n", r.MatchString("Hello world"))

	// \d ---- digit
	r, _ = regexp.Compile(`\d`)
	fmt.Printf("\\d: %v\n", r.MatchString("seven times seven is 49"))

	// \s ---- [\t\n\f\r] (TAB SPACE CR LF)
	r, _ = regexp.Compile(`\s`)
	fmt.Printf("\\s: %v\n", r.MatchString("/home/bill/My Documents"))

	// \D:not digit \S:not whielspace \W:not word-char
	r, _ = regexp.Compile(`\W`)
	fmt.Printf("\\W: %v\n", r.MatchString("555-shoe"))
	fmt.Printf("\\W: %v\n", r.MatchString("555shoe"))

	// Output:
	// \H: true
	// \d: true
	// \s: true
	// \W: true
	// \W: false
}

func ExampleRepeat() {
	//simple repeation
	s := "Key=Value"
	r, _ := regexp.Compile(`\w+=\w+`)
	fmt.Printf("%v\n", r.FindAllString(s, -1))

	s = "Key = Value"
	r, _ = regexp.Compile(`\w+\s*=\s*\w+`)
	fmt.Printf("%v\n", r.FindAllString(s, -1))

	// Output:
	// [Key=Value]
	// [Key = Value]
}

func ExampleUnicode() {
	//unicode classes
	r, _ := regexp.Compile(`\p{Greek}`)
	if r.MatchString("This is all Γςεεκ to me.") == true {
		fmt.Printf("match ") // Will print 'Match'
	} else {
		fmt.Printf("No match ")
	}

	// Output:
	// match
}

func ExampleGroup() {
	//Groups
	r, _ := regexp.Compile(`(.)at`)
	fmt.Printf("groups: %v\n", r.FindAllStringSubmatch("the cat sat one the mat", -1))

	// Output:
	// groups: [[cat c] [sat s] [mat m]]
}
