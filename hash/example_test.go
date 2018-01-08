package hash_test

import (
	"fmt"
	"zshanjun/ec_startup/hash"
)

func ExampleBase64Encode() {
	fmt.Println(hash.Base64Encode("hello"))
	// Output: aGVsbG8=
}

func ExampleBase64Decode() {
	decode, _ := hash.Base64Decode(hash.Base64Encode("hello"))
	fmt.Println(decode)
	// Output: hello
}

func ExampleMd5() {
	fmt.Println(hash.Md5("hello"))
	// Output: 5d41402abc4b2a76b9719d911017c592
}

func ExampleSha1() {
	fmt.Println(hash.Sha1("hello"))
	// Output: aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}

func ExampleCrc32() {
	fmt.Println(hash.Crc32("18124004512"))
	// Output: 2918270904
}
