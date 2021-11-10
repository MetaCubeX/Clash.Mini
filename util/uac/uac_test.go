package uac

import (
	"os"
	"testing"
)

func TestAmAdmin(t *testing.T) {
	if value := os.Getenv("GIT_BRANCH"); len(value) > 0 {
		return
	}
	if AmAdmin {
		t.Fail()
	}
}
