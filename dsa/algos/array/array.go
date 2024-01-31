package array

import (
	"Gapp/dsa/util"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Salgo byte

const (
	Bubble Salgo = iota
	Selection
	Insertion
	Heap
	Quick
	Merge
	Counting
	Radix
	Bucket
)

func BinarySearch(input []int, key int) (bool, int) {
	l, r := 0, len(input)
	for l < r {
		m := l + ((r - l) / 2)
		if key > input[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	if input[l] == key {
		return true, l
	} else {
		return false, -1
	}
}

func binarySearchR(input []int, left int, right int, key int) bool {
	if left < right {
		m := left + ((right - left) / 2)
		if input[m] == key {
			return true
		} else if key > input[m] {
			binarySearchR(input, m+1, right, key)
		} else {
			binarySearchR(input, left, m, key)
		}
	}
	return false
}

// arrange numbers in ai<aj  where i<j
func Sort(input []int, algo Salgo) {
	length := len(input)
	switch algo {
	case Bubble:
		for i := 0; i < length; i++ {
			for j := 0; j < length-i-1; j++ {
				if input[j] > input[j+1] {
					Swap(input, j, j+1)
				}
			}
		}
		break
	case Selection:
		for i := 0; i < length; i++ {
			minIdx := i
			for j := i + 1; j < length; j++ {
				if input[minIdx] > input[j] {
					minIdx = j
				}
			}
			if minIdx != i {
				Swap(input, minIdx, i)
			}
		}
		break
	case Insertion:
		// {10,30,40,50,70,80,90}
		// {10,30,30,40,40,50,70}
		for i := 1; i < len(input); i++ {
			j := i - 1
			for j >= 0 && input[i] < input[j] {
				j--
			}
			j++
			//Copy empty j by moving elements to
			if j != i {
				temp := input[i]
				for k := i - 1; k >= j; k-- {
					input[k+1] = input[k]
				}
				input[j] = temp
			}
		}
		break
	case Heap:
		break

	case Quick:
		QuickSort(input, 0, (len(input) - 1))
		break

	case Merge:
		MergSort(input, 0, length-1)
		break

	}

}

func QuickSort(input []int, left, right int) {
	if left < right {
		pivot := quickPartition(input, left, right)

		fmt.Printf("Left=[%v, %v], Right=[%v,%v] \n", left, pivot-1, pivot+1, right)
		QuickSort(input, left, pivot-1)
		QuickSort(input, pivot+1, right)
	}

}

func quickPartition(input []int, left, right int) int {
	pivot := input[right]
	i := left - 1
	for j := left; j < right; j++ {
		if input[j] < pivot {
			i = i + 1
			Swap(input, i, j)
		}
	}
	Swap(input, i+1, right)
	return (i + 1)
}

func Swap(input []int, i, j int) {
	temp := input[i]
	input[i] = input[j]
	input[j] = temp
}

// Main function that sorts arr[l..r] using
// merge()
func MergSort(arr []int, l, r int) {
	if l < r {
		// Find the middle point
		m := l + (r-l)/2

		// Sort first and second halves
		MergSort(arr, l, m)
		MergSort(arr, m+1, r)

		// Merge the sorted halves
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	// Find sizes of two subarrays to be merged
	n1 := m - l + 1
	n2 := r - m

	/* Create temp arrays */
	var L []int = make([]int, n1)
	var R []int = make([]int, n2)

	/*Copy data to temp arrays*/
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	/* Merge the temp arrays */

	// Initial indexes of first and second subarrays
	i := 0
	j := 0

	// Initial index of merged subarray array
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	/* Copy remaining elements of L[] if any */
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	/* Copy remaining elements of R[] if any */
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func Subset(arr []int) []string {
	nofSubsets := int(math.Pow(2, float64(len(arr))))
	collector := util.StringCollector(nofSubsets)
	for idx := 0; idx < nofSubsets; idx++ {
		i := -1
		var sb strings.Builder
		sb.WriteString("{")
		num := idx
		for num > 0 {
			bit := num & 1
			i += 1
			if bit == 1 {
				fmt.Fprintf(&sb, "%d", arr[i])
			}
			num = num >> 1
		}
		sb.WriteString("}")
		collector.Append(sb.String())
	}

	return collector.Elements
}

// Sets containing element and not containg element.
// // 1,2,3 -> 1,2,3
//     and  2,3

func SubsetBackTracking(arr []int, l, r int) {
	if len(arr) == 1 {
		util.Printable(arr, l, r)
	} else if len(arr) > 1 {
		util.Printable(arr, l, r)
		subsize := r - l
		for i := l + 1; i < len(arr); i++ {
			include := make([]int, 0, subsize)
			for j := l; j < len(arr); j++ {
				if i != j {
					include = append(include, arr[j])
				}
			}
			//	fmt.Printf("include =%v", include)

			SubsetBackTracking(include, 0, len(include)-1)
		}

		exclude := make([]int, subsize, subsize)
		copy(exclude, arr[l+1:])
		if len(exclude) > 1 {
			//fmt.Printf("Exclude =%v", exclude)
			SubsetBackTracking(exclude, 0, len(exclude)-1)
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
			util.Printable(selection[:], 0, 1)
		}
	}
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

// Range queries can be solved by Prefix sum array or Segment Trees and Binary Indexed trees.

func RangeQueriesSum(input []int, queries [][2]int) []int {
	sumArray := make([]int, len(input))
	output := make([]int, len(queries))
	preSum := 0
	for i, val := range input {
		sumArray[i] = preSum + val
	}
	for j := range queries {
		from := queries[j][0]
		to := queries[j][1]

		output[j] = sumArray[to] - sumArray[from-1]
	}
	return output
}

// Range queries  Building segment tree.
func RangeQueriesSegmentTree(input []int, queries [][2]int) []int {
	length := len(input)
	segmentTreeSize := length * (length - 1)
	segmentTree := make([]int, segmentTreeSize)
	output := make([]int, len(queries))

	buildSegmentTree(input, segmentTree, 0, 0, len(input))

	for j := range queries {
		from := queries[j][0]
		to := queries[j][1]
		output[j] = QuerySegemnt(segmentTree, from, to)
	}
	return output
}

func buildSegmentTree(input []int, segmentTree []int, sidx, left, right int) int {
	if left == right {
		segmentTree[left] = input[sidx]
		return input[left]
	} else {
		mid := getMid(left, right)
		leftSum := buildSegmentTree(input, segmentTree, sidx*2+1, left, mid)
		rightSum := buildSegmentTree(input, segmentTree, sidx*2+2, mid+1, right)
		segmentTree[sidx] = leftSum + rightSum
		return segmentTree[sidx]
	}
}

func QuerySegemnt(segement []int, from, to int) int {
	return 0

}

func getMid(l, r int) int {
	return l + (r-l)/2
}

/*
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
*/
func tripleSum(nums []int) {

	sort.Ints(nums)

}

/*

	 {1, 2, 3, 4, 5, 6, 7}
k=3, {5, 6, 7, 1, 2, 3, 4}
k=2, {6, 7, 1, 2, 3, 4, 5}
k=1, {7, 1, 2, 3, 4, 5, 6}
*/

func Rotation(nums []int, k int) {
	l := 0
	r := len(nums) - 1 - k
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l = len(nums) - k
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l = 0
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

}

// {2,3,1,1,4}
func Jump(nums []int) bool {
	lidx := len(nums) - 1
	if lidx == -1 {
		return false
	}
	if lidx == 0 {
		return true
	}

	temp := make([]bool, len(nums))
	END := lidx
	for i := lidx - 1; i >= 0; i-- {
		maxJupms := nums[i]
		for k := maxJupms; k > 0; k-- {
			if i+k <= END {
				if i+k == END || temp[i+k] == true {
					temp[i] = true
				}
			}
		}
	}
	return temp[0]
}

func Dummy(nums1 []int, nums2 []int) float64 {
	last1, last2, sum := 0, 0, 0
	len1, len2 := len(nums1), len(nums2)
	if len1 > 0 && len2 > 0 {
		last1 = nums1[len1-1]
		last2 = nums2[len2-1]
		if last1 > last2 {
			sum = (last1 * (last1 + 1)) / 2
		} else {
			sum = (last2 * (last2 + 1)) / 2
		}
		return float64(sum) / float64(len1+len2)
	} else if len(nums1) > 0 {
		if (len1) == 1 {
			return float64(nums1[0])
		}
		last1 = nums1[len1-1]
		sum = (last1 * (last1 + 1)) / 2
		return float64(sum) / float64(len1)
	} else if len(nums2) > 0 {
		if (len2) == 1 {
			return float64(nums2[0])
		}
		last2 = nums2[len2-1]
		sum = (last2 * (last2 + 1)) / 2
		return float64(sum) / float64(len2)
	}
	return float64(0)
}
