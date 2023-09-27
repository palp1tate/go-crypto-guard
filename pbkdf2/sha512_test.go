package pwd

import "testing"

func TestSHA512(t *testing.T) {
	pwd := "12345678"
	encodedPwd, err := GenSHA512(pwd, 12, 16, 100)
	if err != nil {
		t.Error(err)
	}
	// pwd = "1234567"
	t.Log("encodedPwd:", encodedPwd)
	if ok, err := VerifySHA512(pwd, encodedPwd); !ok {
		t.Error("VerifySHA512 failed", err)
	} else {
		t.Log("VerifySHA512 success")
	}
}
