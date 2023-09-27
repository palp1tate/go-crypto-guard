package pwd

import "testing"

func TestBlake2b(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := Blake2b(pwd)
	if err != nil {
		t.Error(err)
	}
	t.Log(encodedPwd)
	// pwd = "1234567"
	if ok, err := VerifyBlake2b(pwd, encodedPwd); !ok {
		t.Error("VerifyBlake2b failed", err)
	} else {
		t.Log("VerifyBlake2b success")
	}
}
