package string

import "testing"

func TestStartsWith(t *testing.T) {
	s := `Clash.Mini`
	starts := `Clash.Mini`

	if !StartsWith(s, starts) {
		t.Fail()
	}

	s = `Clash.Mini`
	starts = `Clash`

	if !StartsWith(s, starts) {
		t.Fail()
	}

	s = `Clash`
	starts = `Clash.Mini`

	if StartsWith(s, starts) {
		t.Fail()
	}

	s = `Clash.Mini`
	starts = `abcdefg`

	if StartsWith(s, starts) {
		t.Fail()
	}

	s = `ABCDEFG`
	starts = `Clash`

	if StartsWith(s, starts) {
		t.Fail()
	}
}

func TestEndsWith(t *testing.T) {
	s := `Clash.Mini`
	ends := `Clash.Mini`

	if !EndsWith(s, ends) {
		t.Fail()
	}

	s = `Clash.Mini`
	ends = `Mini`

	if !EndsWith(s, ends) {
		t.Fail()
	}

	s = `Mini`
	ends = `Clash.Mini`

	if EndsWith(s, ends) {
		t.Fail()
	}

	s = `Clash.Mini`
	ends = `abcdefg`

	if EndsWith(s, ends) {
		t.Fail()
	}

	s = `ABCDEFG`
	ends = `Clash`

	if EndsWith(s, ends) {
		t.Fail()
	}
}
