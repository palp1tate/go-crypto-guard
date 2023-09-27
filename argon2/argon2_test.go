package pwd

import "testing"

func TestArgon2(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := GenArgon2(pwd, 12, 16, 100)
	if err != nil {
		t.Error(err)
	}
	t.Log(encodedPwd)
	// pwd = "1234567"
	if ok, err := VerifyArgon2(pwd, encodedPwd); !ok {
		t.Error("VerifyScrypt failed", err)
	} else {
		t.Log("VerifyScrypt success")
	}
}
