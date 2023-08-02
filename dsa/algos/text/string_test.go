package text

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	inputs := []string{"", "v", "ve", "venugopal"}
	outputs := []string{"", "v", "ev", "lapogunev"}
	for i := 0; i < len(inputs); i++ {
		reverse := Reverse(inputs[i], 0, len(inputs[i])-1)
		if reverse != outputs[i] {
			fmt.Println(outputs[i], reverse)
			t.Fail()
		}
	}
}

func TestPalindrome(t *testing.T) {
	var in string = "venunev"
	isPalindrome := IsPalidrome(in, 0, len(in)-1)
	if !isPalindrome {
		t.Log("Not a palindrome")
		t.Fail()
	} else {
		t.Log("Is palindrome")
	}
}
