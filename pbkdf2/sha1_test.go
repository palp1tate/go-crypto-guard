package pwd

import "testing"

func TestSHA1(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := SHA1(pwd, 12, 32, 13)
	if err != nil {
		t.Error(err)
	}
	// pwd = "1234567"
	if ok, err := VerifySHA1(pwd, encodedPwd); !ok {
		t.Error("VerifySHA1 failed", err)
	} else {
		t.Log("VerifySHA1 success")
	}
}
