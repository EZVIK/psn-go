package authenticate

import "testing"

func Test_exchangeForAuthToken(t *testing.T) {

	code, err := ExchangeForCode(NPSSO + "123")
	if err != nil {
		t.Errorf("ExchangeForCode error: %v", err)
		return
	}

	accessToken, refreshToken, err := ExchangeForAuthToken(code)
	if err != nil {
		t.Errorf("ExchangeForAuthToken error: %v", err)
		return
	}

	t.Logf("AccessToken: %s,\n RefreshToken: %s", accessToken, refreshToken)
}
