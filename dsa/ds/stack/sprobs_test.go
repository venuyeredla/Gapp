package stack

import "testing"

func TestItoP(t *testing.T) {
	//ItoP("a+b*c+d")

	ItoP("a+b*(c^d-e)^(f+g*h)-i ")
}

func TestBalanced(t *testing.T) {
	balanced := IsBalanced("[()]{}{[()()]()}")
	if !balanced {
		t.Fail()
	}

}
