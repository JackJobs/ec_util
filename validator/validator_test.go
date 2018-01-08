//参考https://github.com/caixw/lib.go/blob/master/validation/validator/regexp.go
package validator

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	if i := IsEmail("zshanjun@qq.com"); i != true {
		t.Error("fail")
	}
	if i := IsEmail("kdkkdk@dd"); i != false {
		t.Error("fail")
	}
}

func TestIsMobile(t *testing.T) {
	if i := IsMobile("18124004512"); i != true {
		t.Error("fail")
	}
	if i := IsMobile("01254621"); i != false {
		t.Error("fail")
	}
}

func TestIsIp4(t *testing.T) {
	if i := IsIp4("192.168.1.1"); i != true {
		t.Error("fail")
	}
	if i := IsIp4("259.123.4.5"); i != false {
		t.Error("fail")
	}
}

func TestIsUrl(t *testing.T) {
	if i := IsUrl("https://www.workec.com"); i != true {
		t.Error("fail")
	}
	if i := IsUrl("www.workec.com"); i != true {
		t.Error("fail")
	}
	if i := IsUrl("kkdkkdi"); i != false {
		t.Error("fail")
	}
}


