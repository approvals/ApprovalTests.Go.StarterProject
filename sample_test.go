package starterpackage

import (
	"fmt"
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

func TestJsonApproval(t *testing.T) {
	person := Person{
		"John Galt", 100,
	}
	approvals.VerifyJSONStruct(t, person)
}

func TestCombinationsApproval(t *testing.T) {
	arr := []int{10, 20, 30, 40}
	names := []string{"john", "paul", "mary"}
	approvals.VerifyAllCombinationsFor2(t, "File Names", func(a interface{}, b interface{}) string { return fmt.Sprintf("%v_%v.txt", a, b) }, names, arr)
}
