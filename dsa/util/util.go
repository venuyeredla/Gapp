package util

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func AssertEquals(expected, actual []int, log bool) (result bool, message string) {
	if log {
		Printable(expected, 0, len(expected)-1)
		Printable(actual, 0, len(expected)-1)
	}
	if len(expected) == len(actual) {
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				return false, "Failed at index - " + fmt.Sprint(i) + " Expected = " + fmt.Sprint(expected[i]) + " Actual = " + fmt.Sprint(actual[i])
			}
		}
		return true, ""
	} else {
		return false, "Array lengths are unequal"
	}
}

func Printable(arr []int, l, r int) string {
	if l <= r {
		var sb strings.Builder
		sb.WriteString("{")
		for k := l; k <= r; k++ {
			fmt.Fprintf(&sb, "%d", arr[k])
			if k < r {
				sb.WriteString(",")
			}
		}
		sb.WriteString("}")
		s := sb.String()
		fmt.Println(s)
		return s
	}
	return ""
}

func GenArray(size int, max int) []int {
	// balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}  // static arry with intialization of values.
	//var generated [5]int // fixed array size declaration
	var generated []int //If size don't mentined it will become slice. Before using slice need to intilaize.
	rand.Seed(time.Now().UnixMilli())
	generated = make([]int, size)
	for i := 0; i < size; i++ {
		generated[i] = rand.Intn(20)
	}
	return generated
}

func GetMatrix(rows, cols int) [][]int {
	sm := make([][]int, rows)
	for i := 0; i < rows; i++ {
		sm[i] = make([]int, cols)
	}
	return sm
}

func Minimum(i, j, k int) int {
	return int(math.Min(float64(math.Min(float64(i), float64(j))), float64(k)))
}

func MaxOf(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
