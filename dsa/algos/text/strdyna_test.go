package text

import (
	"fmt"
	"testing"
)

func TestSubArray(t *testing.T) {
	str := "abcd"
	fmt.Println("Using iteration : ")
	SubStrings(str, 0, len(str), len(str))
	/*
	   fmt.Println("Using bottom up recursion : ")
	   SubStringsR(str, 0, 0, 4)
	*/
}

func TestSubstrs(t *testing.T) {
	//input := []string{"a", "b", "c"}
	input := "aaaabbaa"
	Substrs(input, 0, 0, len(input))
}

func TestEditDistance(t *testing.T) {
	distance := ediDistance("saturday", "sunday")
	dd := ediDistanceD("saturday", "sunday")

	if distance != 3 && dd != 3 {
		t.Logf("Distance is not correct")
		t.Fail()
	}
}

func TestLcs(t *testing.T) {
	lcs := Lcs("aggtab", "gxtxayb")
	lcsd := LcsD("aggtab", "gxtxayb")

	if lcs != 4 && lcsd != 4 {
		t.Logf("Distance is not correct")
		t.Fail()
	}
}

func TestLrs(t *testing.T) {
	input := "aabebcdd"
	l := len(input) - 1
	lcs := Lrs(input, l, l)
	if lcs != 3 {
		t.Logf("Distance is not correct")
		t.Fail()
	}
}

// bbabcbcab => babcbab
func TestLps(t *testing.T) {
	input := "bbabcbcab"
	lcs := Lps(input)
	LpsD("bbabcbcab")
	if lcs != 7 {
		t.Logf("Distance is not correct")
		t.Fail()
	}
}

func TestWordBreak(t *testing.T) {

	words := []string{"leet", "code"}
	actual := wordBreak(words, "leetcode")
	expected := true

	if actual != expected {
		t.Errorf("Wrong logic")
	}

}
