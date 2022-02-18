package starterpackage

import (
	approvals "github.com/approvals/go-approval-tests"
	"testing"
)

func TestNormalTest(t *testing.T) {
	if 5 != 5 {
		t.Fatal("5 != 5")
	}
}

func TestBasicApproval(t *testing.T) {
	approvals.VerifyString(t, "Hello Approvals")
}
