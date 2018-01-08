package crypto_test

import (
	"fmt"
	cryptoutil "zshanjun/ec_startup/crypto"
	"zshanjun/ec_startup/hash"
)

func ExampleDes() {
	key := []byte("vvFormvv")
	result, _ := cryptoutil.DesEncrypt([]byte("122611"), key, "ecb")
	fmt.Println(hash.Base64Encode(string(result)))
	src, _ := cryptoutil.DesDecrypt(result, key, "ecb")
	fmt.Println(string(src))
	// Output:
	// 3YcMXdt6ORY=
	// 122611
}

