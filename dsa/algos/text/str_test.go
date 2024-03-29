package text

import (
	"Gapp/dsa/algos/maths"
	"Gapp/dsa/util"
	"fmt"
	"testing"
)

type TestIO struct {
	Input  string
	Output string
}

func TestPermuations(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	size := maths.Factorial(len(arr))
	collector := util.StringCollector(size)
	Permuations(arr, 0, len(arr), collector)
	for _, val := range collector.Elements {
		fmt.Println(val)
	}
}

func TestSubSequnces(t *testing.T) {
	subseq := SubSeqences("bbabcbcab") //bbabcbcab     abc
	for _, s := range subseq {
		if IsPalindrome(s) {
			fmt.Printf("%v - %v\n", s, len(s))
		}

	}
}

func TestReverse(t *testing.T) {
	io := make([]TestIO, 0, 4)
	io = append(io, TestIO{"", ""})
	io = append(io, TestIO{"v", "v"})
	io = append(io, TestIO{"ve", "ev"})
	io = append(io, TestIO{"venugopal", "lapogunev"})
	for _, val := range io {
		reverse := ""
		if reverse != val.Output {
			fmt.Println(val.Output, reverse)
			t.Fail()
		}
	}
}

func TestString(t *testing.T) {
	MultiStr("0033", "22")
}

func TestRToD(t *testing.T) {
	num := RomanDecimal("MCMXCIV")
	if num == 1994 {
		t.Log("Not a valid decimal")
		t.Fail()
	}
}
