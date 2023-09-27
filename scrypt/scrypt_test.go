package pwd

import "testing"

func TestScrypt(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := GenScrypt(pwd, 16, 32)
	if err != nil {
		t.Error(err)
	}
	t.Log("encodedPwd:", encodedPwd)
	if ok, err := VerifyScrypt(pwd, encodedPwd); !ok {
		t.Error("VerifyScrypt failed", err)
	} else {
		t.Log("VerifyScrypt success")
	}
}
