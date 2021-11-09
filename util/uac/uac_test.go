package uac

import "testing"

func TestAmAdmin(t *testing.T) {
	if AmAdmin {
		t.Fail()
	}
}