package text

import (
	"Gapp/dsa/util"
	"fmt"
	"strconv"
	"strings"
)

// Backtracking algorithm
func Permuations(arr []int, left, right int, collector *util.Collector) {
	if left == right {
		var sb strings.Builder
		for k := 0; k <= right-1; k++ {
			fmt.Fprintf(&sb, "%d", arr[k])
		}
		collector.Append(sb.String())
	} else {
		for i := left; i < right; i++ {
			SwapA(arr, i, left)
			Permuations(arr, left+1, right, collector)
			SwapA(arr, left, i) //Backtracking step
		}
	}
}

func SwapA(input []int, i, j int) {
	temp := input[i]
	input[i] = input[j]
	input[j] = temp
}

/*
Input : abc
Output : a, ab, abc, ac   b, bc, c

	inclusion and eclusion principle. Backtracking and recurison
*/
func SubSeqences(str string) []string {
	collector := util.StringCollector(8)
	subSeq(str, "", 0, len(str), collector)
	return collector.Elements
}

func subSeq(str, sub string, left, right int, collector *util.Collector) {
	if str == "" || left == right {
		return
	}
	//for i := left; i < right; i++ {
	newSub := sub + string(str[left])
	collector.Append(newSub)
	// Swap(&str, i, left)
	subSeq(str, newSub, left+1, right, collector)
	subSeq(str, sub, left+1, right, collector)
}

/*
if size

	even -> character frequency should be even
	odd -> one character should have odd freqencies.
*/
func IsPalindrome(str string, l, r int) bool {
	if l >= r {
		return true
	} else if str[l:l+1] != str[r:r+1] {
		return false
	} else {
		return IsPalindrome(str, l+1, r-1)
	}
}

func Palindrome(A string) bool {
	l := 0
	r := len(A) - 1
	for l <= r {
		if A[l] != A[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func TestType() bool {
	var a uint16 = 13
	var b uint16 = 13
	return a-b < 0
}

func PMatchNaive(text, pattern string) int {
	index := -1
	matched := true
	for tIdx := 0; tIdx <= len(text)-len(pattern); tIdx++ {
		index = tIdx
		matched = true
		for pIdx := 0; pIdx < len(pattern); pIdx++ {
			temp := pIdx + tIdx
			if pattern[pIdx:pIdx+1] != text[temp:temp+1] {
				matched = false
				index = -1
				break
			}
		}
		if matched == true {
			break
		}
	}
	return index
}

func searchKMP(txt, pat string) int {
	//int[] buildLps = this.buildLps(pat);
	for i := 0; i < len(txt); i++ {

	}

	return 0
}

func buildLps(pat string) []int {
	lps := make([]int, len(pat))
	i := 0
	j := 1
	lps[i] = 0
	for j < len(pat) {
		if pat[i] == pat[j] {
			lps[j] = i + 1
			i++
			j++
		} else {
			if i != 0 {
				i = lps[i-1]
			} else {
				lps[j] = i
				j++
			}
		}
	}
	return lps
}

func validIp(ip string) bool {
	if len(ip) > 0 {
		return false
	}
	ipparts := strings.Split(ip, "\\.")

	if len(ipparts) != 4 {
		return false
	}
	for _, value := range ipparts {
		num, error := strconv.Atoi(value)
		if error == nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func makeIP(str string) string {
	/*if(str==null || str=="" || str.length()<=3 || str.length()>=13) {
		throw new IllegalArgumentException("IP can't be make with : "+str);
	} */
	//n := len(str) // n=4*a+r;
	//a := n / 4
	// re := n % 1
	// l := 0
	// r := 0
	//	var lr [][]int=make() int[4][2];

	return ""
}

var RIMap = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

/*

 I, II, III, IV, V , VI , VII, VIII, IX, X, XI

 1994 - MCMXCIV  =? 2194
*/

func RomanDecimal(str string) int {
	integer := 0
	PRE := ""
	for i := len(str) - 1; i >= 0; i-- {
		symbol := string(str[i])
		val, _ := RIMap[symbol]
		if symbol == "I" && (PRE == "V" || PRE == "X") {
			integer = integer - val
		} else if symbol == "X" && (PRE == "L" || PRE == "C") {
			integer = integer - val
		} else if symbol == "C" && (PRE == "D" || PRE == "M") {
			integer = integer - val
		} else {
			integer = integer + val
		}
		PRE = symbol
	}
	return integer
}

func MultiStr(a, b string) {
	for i := len(b) - 1; i >= 0; i-- {
		bc := b[i : i+1]
		for j := len(a) - 1; j >= 0; j-- {
			ac := a[j : j+1]
			aa, _ := strconv.Atoi(ac)
			bb, _ := strconv.Atoi(bc)
			result := aa * bb
			fmt.Printf("%v", result)
		}
		fmt.Println("")
	}
}

// Size=(n *(n+1))/2  ==> l=r=0 means bottom up;
func SubStrings(str string, l, r, n int) {
	if r < n {
		fmt.Println(str[l : r+1])
		SubStrings(str, l, r+1, n)
		if r+1 == n && (l+1) < n {
			SubStrings(str, l+1, l+1, n)
		}
	}
}

var max string
var maxlen = 0

func Substrs(A string, l, r, n int) {
	if r == n {
		return
	}
	if IsPalindrome(A, l, r) {
		size := len(A[l : r+1])
		if size > maxlen {
			maxlen = size
			max = A[l : r+1]
			fmt.Printf("Needed = %v \n", max)

		} else {
			fmt.Printf("Not needed = %v \n", A[l:r+1])

		}
	}

	Substrs(A, l, r+1, n)
	if r+1 == n {
		Substrs(A, l+1, l+1, n)
	}
}

func wordPattern(pattern string, s string) bool {
	if len(s) == 0 || len(pattern) == 0 {
		return false
	}
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}
	lm := make(map[string]string)
	for i, val := range pattern {
		pk := string(val)
		sval := strs[i]
		mval, exist := lm[pk]
		if exist {
			if mval != sval {
				return false
			}
		} else {
			lm[pk] = sval
		}
	}
	return true
}
