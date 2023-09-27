package pwd

import "testing"

func TestMd5(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := Md5(pwd, 12, 32, 13)
	if err != nil {
		t.Error(err)
	}
	// pwd = "1234567"
	if ok, err := VerifyMd5(pwd, encodedPwd); !ok {
		t.Error("VerifyMd5 failed", err)
	} else {
		t.Log("VerifyMd5 success")
	}
}
