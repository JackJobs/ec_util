package hash

import (
	"encoding/base64"
	"crypto/md5"
	"crypto/sha1"
	"hash/crc32"
	"encoding/hex"
)

//与php base64_encode()相同
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

//与php base64_decode()相同
func Base64Decode(src string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(src)
	return string(decode), err
}

//与php md5()相同
func Md5(src string) string {
	res := md5.Sum([]byte(src))
	return hex.EncodeToString(res[:])
}

//与php sha1()相同
func Sha1(src string) string {
	res := sha1.Sum([]byte(src))
	return hex.EncodeToString(res[:])
}

//与php crc32()相同
func Crc32(src string) uint32 {
	return crc32.ChecksumIEEE([]byte(src))
}
