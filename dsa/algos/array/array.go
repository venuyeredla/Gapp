package array

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

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

// Size=(n *(n+1))/2
func RSubArrays(arr []int, l, r, n int) {
	if r == n {
	} else if r < n {
		Printable(arr, l, r)
		RSubArrays(arr, l, r+1, n)
		if r+1 == n && (l+1) < n {
			RSubArrays(arr, l+1, l+1, n)
		}
	}
}

// Backtracking algorithm
func Permuations(arr []int, left, right int) {
	if left == right {
		Printable(arr, 0, right-1)
	} else {
		for i := left; i < right; i++ {
			Swap(arr, i, left)
			Permuations(arr, left+1, right)
			Swap(arr, left, i)
		}
	}
}

// Assumption arr has unique elements {1,2,3}
func combinations(arr []int, to, size int) {
	/* for i := 0; i < len(arr); i++ {
		Swap(arr, to, i)
		PrintArr(arr, 0, to)
		Swap(arr, i, to)
	} */
	selection := make([]int, size) //[2]int{-1, -1}
	for i := 0; i < len(arr); i++ {
		selection[0] = arr[i]
		for j := i + 1; j < len(arr); j++ {
			selection[1] = arr[j]
			Printable(selection[:], 0, 1)
		}
	}
}

func LargestSumContiguous(a []int) {
	//StringJoiner stringJoiner=new StringJoiner(", ");
	Printable(a, 0, len(a)-1)
	sum := 0
	fidx := 0
	toIdx := 0
	maxSum := 0
	for j := 0; j < len(a); j++ {
		sum = sum + a[j]
		if maxSum <= sum {
			maxSum = sum
			toIdx = j
		}
		if sum < 0 {
			sum = 0
			fidx = j
		}
		//stringJoiner.add(sum + "")
	}
	//System.out.print(stringJoiner.toString())
	fmt.Printf(" => {%v,%v = %v} \n", fidx, toIdx, maxSum)
}

// Partion negative and positive numbers
func partion(a []int) {
	pivot := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[pivot] {
			Swap(a, pivot, i)
			pivot++
		}
	}
}

// Print numbers from forward and backward.
func PrintFB(a []int, i int) {
	if len(a) == i {
		return
	} else {
		fmt.Printf("%v ,", a[i])
		PrintFB(a, i+1)
		fmt.Printf("%v ,", a[i])
	}
}

func sum(a []int, i int) int {
	if len(a) == i {
		return a[i]
	} else {
		return a[i] + sum(a, i+1)
	}
}

// {1, 2, 0, 0, 0, 3, 6};
// Method-1 : whenever zero element move all next elements to before then place zero at the end.
// Method-2 : Counting zeros and moving next nonzero elemnt
func MovallZeros(a []int) {
	zeros := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			if zeros > 0 {
				a[i-zeros] = a[i]
				a[i] = 0
			}
		} else {
			zeros++
		}
	}

}

// a[i]=i  a[i]=-1, a[]!=i
// {-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
func Rearrange(a []int) {
	for i := 0; i < len(a); i++ {
		if a[i] != i && a[i] != -1 {
			t := a[i]
			for a[t] != -1 && a[t] != t {
				t2 := a[t]
				a[t] = t
				t = t2
			}
			a[t] = t
			if a[i] != i {
				a[i] = -1
			}

		}
	}
}

func KMostOccur(a []int) {
	freqMap := make(map[int]int)
	for _, v := range a {
		_, present := freqMap[v]
		if present {
			freqMap[v]++
		} else {
			freqMap[v] = 1
		}
	}
	keys := make([]int, len(freqMap))
	for k, v := range freqMap {
		keys = append(keys, k)
		fmt.Printf("%v - %v ,", k, v)
	}
	fmt.Println()
	sort.Ints(keys)

	for _, k := range keys {
		value, exist := freqMap[k]
		if exist {
			fmt.Printf("%v - %v , ", k, value)
		}
	}
	fmt.Println()
	sort.SliceStable(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})
	for _, k := range keys {
		value, exist := freqMap[k]
		if exist {
			fmt.Printf("%v - %v , ", k, value)
		}
	}
	fmt.Println()
}

// Sorted array
func removeDuplicates(input []int) {
	i := 0
	for j := 1; j < len(input); j++ {
		if input[j] != input[i] {
			i = i + 1
			input[i] = input[j]
		}
	}
	i++
	for ; i < len(input); i++ {
		input[i] = -1
	}

	fmt.Printf("%v", input)

}

func BuyAndsell(s []int) {
	n := len(s)
	maxProfit := 0
	maxPrice := s[n-1]
	maxPriceIdx := n - 1
	for j := n - 2; j >= 0; j-- {
		if maxPrice < s[j] {
			maxPrice = s[j]
			maxPriceIdx = j
		}
		local := maxPrice - s[j]
		if maxProfit < local {
			maxProfit = local
		}
	}

	fmt.Println(maxPriceIdx)
}
