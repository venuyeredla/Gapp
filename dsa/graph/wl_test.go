package graph

import (
	"testing"
)

func TestWordLadder(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	restult := ladderLength("hit", "cog", wordList)
	if restult != 5 {
		t.Errorf("Failed Result = %v, Acutal = 5", restult)
	}
}
