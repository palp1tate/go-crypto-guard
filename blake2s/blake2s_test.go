package pwd

import "testing"

func TestBlake2s(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := GenBlake2s(pwd)
	if err != nil {
		t.Error(err)
	}
	t.Log(encodedPwd)
	// pwd = "1234567"
	if ok, err := VerifyBlake2s(pwd, encodedPwd); !ok {
		t.Error("VerifyBlake2s failed", err)
	} else {
		t.Log("VerifyBlake2s success")
	}
}
