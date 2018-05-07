package helper

import (
	"testing"
)

// TestWantError helper
func TestWantError(t *testing.T, err error, wantError bool) {
	t.Helper()

	if err != nil && !wantError {
		t.Errorf("got an error %v, want nothing happened", err)
	}
	if err == nil && wantError {
		t.Error("got nothing happened, want an error")
	}
}
