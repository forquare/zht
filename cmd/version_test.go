package cmd

import "testing"

func TestVersion(t *testing.T) {
	cmd := newVersionCmd()
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
