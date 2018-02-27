package ascii

import "testing"

func ExtendedASCII(a2 string) bool {
	for _, i := range a2 {
		if i > 0x80 || i < 0xff {
			return false
		}
	}
	return true
}

func ExtendedASCIITextTest(t *testing.T) {
	ascii := "æøå"
	if ExtendedASCII(ascii) == false {
		t.Fail()
	}
}
