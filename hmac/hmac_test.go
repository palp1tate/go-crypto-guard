package pwd

import "testing"

func TestHMAC(t *testing.T) {
	pwd := "123456"
	encodedPwd, err := HMAC(pwd, 16)
	if err != nil {
		t.Error(err)
	}
	t.Log(encodedPwd)
	// pwd = "1234567"
	if ok, err := VerifyHMAC(pwd, encodedPwd); !ok {
		t.Error("VerifyHMAC failed", err)
	} else {
		t.Log("VerifyHMAC success")
	}
}
