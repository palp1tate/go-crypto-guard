package pwd

import "testing"

func TestDES(t *testing.T) {
	pwd := "1234567855225322633uby"
	key := "wxy45678"
	encodedPwd, _ := GenDES(pwd, key)
	t.Log(encodedPwd)
	isValid, _ := VerifyDES(pwd, encodedPwd, key)
	if !isValid {
		t.Error("VerifyDES failed")
	}
}
