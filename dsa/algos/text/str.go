package text

import (
	"strconv"
	"strings"
)

func Reverse(str string, l, r int) string {
	if l >= r {
		return str
	} else {
		str = Swap(&str, l, r)
		return Reverse(str, l+1, r-1)
	}
}

func Swap(str *string, l, r int) string {
	bytes := []byte(*str)
	if l > -1 && l < len(*str) && r > -1 && r < len(*str) {
		temp := bytes[l]
		bytes[l] = bytes[r]
		bytes[r] = temp
	}
	return string(bytes)
}

func IsPalidrome(str string, l, r int) bool {
	if l >= r {
		return true
	} else if str[l:l+1] != str[r:r+1] {
		return false
	} else {
		return IsPalidrome(str, l+1, r-1)
	}
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

func rToD(romanNumb string) int {
	integer := 0
	PRE := "I"
	for i := len(romanNumb) - 1; i >= 0; i-- {

		symbol := strconv.Itoa(int(romanNumb[i]))
		val, _ := RIMap[symbol]
		if symbol == "I" && (PRE == "V" || PRE == "X") {
			integer = integer - val
		} else if symbol == "X" && (PRE == "L" || PRE == "C") {
			integer = integer - val
		} else if symbol == "C" && (PRE == "D" || PRE == "M") {
			integer = integer - -val
		} else {
			integer = integer + -val
		}
		PRE = symbol
	}
	return integer
}
