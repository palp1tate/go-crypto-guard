package pwd

import "testing"

func TestRC4(t *testing.T) {
	pwd := "123456rsgzerszgr"
	key := "wxy1sagededawsaehsrgrwg23"
	encodedPwd, err := RC4(pwd, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(encodedPwd)
	isValid, err := VerifyRC4(pwd, encodedPwd, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(isValid)
}
