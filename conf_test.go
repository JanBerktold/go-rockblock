package rockblock

import (
	"testing"
)

func TestRegexTimeAnswer(t *testing.T) {
	if !regTimeAnswer.MatchString("+CCLK:02/05/15,22:10:00") {
		t.FailNow()
	}
	if regTimeAnswer.MatchString("+CCLK:02/5/15,22:10:00") {
		t.FailNow()
	}
}

func TestSessionAnswer(t *testing.T) {
	if !regSesResult.MatchString("+SBDIX: -0, a") {
		t.FailNow()
	}
}
