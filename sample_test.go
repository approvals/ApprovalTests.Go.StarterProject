package starterpackage

import "testing"

func TestNormalTest(t *testing.T) {
	if 5 != 5 {
		t.Fatal("5 != 5")
	}
}
