package text

import (
	"fmt"
	"testing"
)

func TestTextJustify(t *testing.T) {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	for _, val := range TextJustify(words, maxWidth) {
		fmt.Printf("% v  -- %v\n", val, len(val))
	}
}
