package pwd

import "testing"

func TestSHA256(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := GenSHA256(pwd, 12, 32, 13)
	if err != nil {
		t.Error(err)
	}
	// pwd = "1234567"
	if ok, err := VerifySHA256(pwd, encodedPwd); !ok {
		t.Error("VerifySHA256 failed", err)
	} else {
		t.Log("VerifySHA256 success")
	}
}
