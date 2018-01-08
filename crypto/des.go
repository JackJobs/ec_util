package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func DesEncrypt(src, key []byte, mode string) ([]byte, error) {
	switch mode {
	case "ecb":
		return ecbEncrypt(src, key)
	case "cbc":
		return desEncrpt(src, key)
	default:
		return ecbEncrypt(src, key)
	}
}

func DesDecrypt(crypted, key []byte, mode string) ([]byte, error) {
	switch mode {
	case "ecb":
		return ecbDecrypt(crypted, key)
	case "cbc":
		return desDecryp(crypted, key)
	default:
		return ecbDecrypt(crypted, key)
	}
}

func desEncrpt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = pKCS5Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(src))
	blockMode.CryptBlocks(crypted, src)
	return crypted, nil
}

func desDecryp(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	src := make([]byte, len(crypted))
	blockMode.CryptBlocks(src, crypted)
	src = pKCS5UnPadding(src)
	return src, nil
}

func ecbEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	 src = pKCS5Padding(src, bs)
	//src = zeroPadding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func ecbDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	 out = pKCS5UnPadding(out)
	//out = zeroUnPadding(out)
	return out, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(src []byte) []byte {
	return bytes.TrimFunc(src, func(r rune) bool {
		return r == rune(0)
	})
}
