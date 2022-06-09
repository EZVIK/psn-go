package main

import (
	"os"
	"testing"
)

func Test_exchangeForCode(t *testing.T) {

	code, err := ExchangeForCode(NPSSO)
	if err != nil {
		t.Errorf("%v", err)
		os.Exit(1)
	}

	t.Logf("code: %s", code)
}
