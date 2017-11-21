package pswd

import (
	"testing"
)

func Test_Check(t *testing.T) {
	ps, _ := New("hello")

	if !ps.Check("hello") {
		t.Errorf("hello no match")
	}

	ps2, _ := Parse(ps.String())

	if !Equal(ps, ps2) {
		t.Errorf("Encode Decode Fail")
	}

	if !ps2.Check("hello") {
		t.Errorf("Decode not checking true")
	}

	if ps.Check("Goodbye") {
		t.Errorf("Decoded, not checking False")
	}
}
