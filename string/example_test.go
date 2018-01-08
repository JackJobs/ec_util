package string_test

import (
	"fmt"
	stringutil "zshanjun/ec_startup/string"
	"strings"
)

//是否存在某个字符或字串
func ExampleContains() {
	fmt.Println(strings.Contains("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("", ""))
	// Output:
	// false
	// true
	// false
}

//字串出现次数
func ExampleCount() {
	fmt.Println(strings.Count("hello", "h"))
	fmt.Println(strings.Count("hello", ""))
	fmt.Println(strings.Count("fivevev", "vev"))
	// Output:
	// 1
	// 6
	// 1
}

//字符串分割为[]string
func ExampleSplit() {
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar   baz   "))
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	// Output:
	// Fields are: ["foo" "bar" "baz"]
	// ["a" "b" "c"]
	// ["" "man " "plan " "canal panama"]
	// [" " "x" "y" "z"]
	// [""]
}

//字符串是否有某个前缀或后缀
func ExamplePrefix() {
	fmt.Println(strings.HasPrefix("abc", "ab"))
	fmt.Println(strings.HasSuffix("abc", "bc"))
	// Output:
	// true
	// true
}

//字符或字串在字符串中出现的位置
func ExampleIndex() {
	fmt.Printf("%d\n", strings.IndexFunc("studygolang", func(c rune) bool {
		if c > 'u' {
			return true
		}
		return false
	}))
	// Output: 4

}

//字符串join操作
func ExampleJoin() {
	fmt.Println(strings.Join([]string{"abc", "ddd"}, "&"))
	// Output: abc&ddd
}

//字符串重复几次
func ExampleRepeat() {
	fmt.Println("ba" + strings.Repeat("na", 2))
	// Output: banana
}

//字符串替换
func ExampleReplace() {
	fmt.Println(strings.Replace("old old old", "old", "new", 2))
	fmt.Println(strings.Replace("old old old", "old", "new", -1))

	//Replacer类型
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>Html</b>!"))

	// Output:
	// new new old
	// new new new
	// This is &lt;b&gt;Html&lt;/b&gt;!
}

// strings.Index的UTF-8版本
func ExampleUtf8Index() {
	fmt.Println(stringutil.Utf8Index("Go语言中文网", "中文"))
	// Output: 4
}
