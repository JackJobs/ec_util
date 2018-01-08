package validator

import "regexp"

const (
	emailStr = `[\w.-]+@[\w_-]+\w{1,}[\.\w-]+`
	//ec_be
	mobileStr = `1[3|4|5|6|7|8|9][0-9]\d{8}`
	ip4Str = `((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`
	//ec_be（有问题）
	urlStr = `((https?|ftp|file):\/\/)?[-A-Za-z0-9+&@#\/%?=~_|!:,.;]+[-A-Za-z0-9+&@#\/%=~_|]`
)

func regexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}

var (
	email = regexpCompile(emailStr)
	mobile = regexpCompile(mobileStr)
	ip4 = regexpCompile(ip4Str)
	url = regexpCompile(urlStr)
)

func isMathch(exp *regexp.Regexp, val interface{}) bool {
	switch v := val.(type) {
	case []rune:
		return exp.MatchString(string(v))
	case []byte:
		return exp.Match(v)
	case string:
		return exp.MatchString(v)
	default:
		return false
	}
}

func IsMobile(val interface{}) bool {
	return isMathch(mobile, val)
}

func IsEmail(val interface{}) bool {
	return isMathch(email, val)
}

func IsIp4(val interface{}) bool {
	return isMathch(ip4, val)
}

func IsUrl(val interface{}) bool {
	return isMathch(url, val)
}
