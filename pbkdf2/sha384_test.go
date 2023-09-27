package pwd

import "testing"

func TestSHA384(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := SHA384(pwd, 12, 32, 13)
	if err != nil {
		t.Error(err)
	}
	// pwd = "1234567"
	if ok, err := VerifySHA384(pwd, encodedPwd); !ok {
		t.Error("VerifySHA384 failed", err)
	} else {
		t.Log("VerifySHA384 success")
	}
}
