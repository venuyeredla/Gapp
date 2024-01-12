package text

import (
	"fmt"
	"testing"
)

func TestTextJustify(t *testing.T) {

	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16

	for _, val := range TextJustify(words, maxWidth) {
		fmt.Println(val)
	}
}
