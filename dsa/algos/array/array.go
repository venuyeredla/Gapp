package array

import (
	heap "Gapp/dsa/ds/priority"
	"Gapp/dsa/ds/types"
	"Gapp/dsa/utils"
	"fmt"
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

// arrange numbers in ai<aj  where i<j
func Sort(input []int, algo Salgo) {
	length := len(input)
	switch algo {
	case Bubble:
		for i := 0; i < length; i++ {
			for j := 0; j < length-i-1; j++ {
				if input[j] > input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
				}
			}
		}

	case Selection:
		for i := 0; i < length; i++ {
			minIdx := i
			for j := i + 1; j < length; j++ {
				if input[minIdx] > input[j] {
					minIdx = j
				}
			}
			if minIdx != i {
				input[minIdx], input[i] = input[i], input[minIdx]
			}
		}

	case Insertion:
		for i := 1; i < len(input); i++ {
			j := i - 1
			temp := input[i]
			for j >= 0 && temp < input[j] {
				input[j+1] = input[j]
				j--
			}
			j++
			if j != i {
				input[j] = temp
			}

		}
	case Heap:
		fmt.Println("Need to be implmented")

	case Quick:
		QuickSort(input, 0, len(input)-1)

	case Merge:
		MergSort(input, 0, length-1)
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
			input[i], input[j] = input[j], input[i]
		}
	}
	input[i+1], input[right] = input[right], input[i+1]
	return (i + 1)
}

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
	//var L []int = make([]int, n1)
	LT := arr[l : m+1]
	RT := arr[m+1 : r+1]
	L := make([]int, len(LT))
	R := make([]int, len(RT))
	copy(L, LT[:])
	copy(R, RT[:])
	i, j := 0, 0
	k := l
	for i < len(L) && j < len(R) {
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
	for i < len(L) {
		arr[k] = L[i]
		i++
		k++
	}

	/* Copy remaining elements of R[] if any */
	for j < len(R) {
		arr[k] = R[j]
		j++
		k++
	}
}

// Since array sorted we can use binary search
func findKeyRotated(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			return m
		} else if target >= nums[l] {
			if target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}

		}
	}
	return -1
}

func PowerSet(arr []int, nofSubsets int, collector *utils.Collector) {
	for idx := 0; idx < nofSubsets; idx++ {
		i := -1
		var sb strings.Builder
		sb.WriteString("{")
		num := idx
		for num > 0 {
			i += 1
			if num&1 == 1 {
				fmt.Fprintf(&sb, "%d", arr[i])
			}
			num = num >> 1
		}
		sb.WriteString("}")
		collector.Append(sb.String())
	}

}

// Sets containing element and not containg element.
// // 1,2,3 -> 1,2,3
//     and  2,3

func PowerSetByBacktrack(arr []int, collector *utils.Collector, l, r int) {
	if len(arr) == 1 {
		collector.Append(utils.Printable(arr, l, r))
	} else if len(arr) > 1 {
		collector.Append(utils.Printable(arr, l, r))
		subsize := r - l
		for i := l + 1; i < len(arr); i++ {
			include := make([]int, 0, subsize)
			for j := l; j < len(arr); j++ {
				if i != j {
					include = append(include, arr[j])
				}
			}
			//	fmt.Printf("include =%v", include)
			PowerSetByBacktrack(include, collector, 0, len(include)-1)
		}

		exclude := make([]int, subsize)
		copy(exclude, arr[l+1:])
		if len(exclude) > 1 {
			//fmt.Printf("Exclude =%v", exclude)
			PowerSetByBacktrack(exclude, collector, 0, len(exclude)-1)
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
			utils.Printable(selection[:], 0, 1)
		}
	}
}

// {1, 2, 0, 0, 0, 3, 6};
// Method-1 : whenever zero element move all next elements to before then place zero at the end.
// Method-2 : Counting zeros and moving next nonzero elemnt
func MovallZeros(a []int) {
	widx := -1
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			widx += 1
			a[widx] = a[i]
		}
	}
	for widx += 1; widx < len(a); widx++ {
		a[widx] = 0
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

/*
Approach 1 : Frequency map => Sort By Value => then pick top K elements.
Approach 2 : using Priority queue.
*/
func KMostOccurance(a []int, k int) []int {

	freqMap := make(map[int]int)
	for _, v := range a {
		_, present := freqMap[v]
		if present {
			freqMap[v]++
		} else {
			freqMap[v] = 1
		}
	}

	pq := heap.NewMinHeap(10)

	for k, v := range freqMap {
		value := types.Int32(k)
		pq.Push(v, value)
	}
	/*
		sort.Ints(keys)
		sort.SliceStable(keys, func(i, j int) bool {
			return freqMap[keys[i]] > freqMap[keys[j]]
		})
	*/
	result := make([]int, k)

	return result
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
}

func BuyAndsell(s []int) int {
	n := len(s)
	maxProfit := 0
	maxPrice := s[n-1]
	for j := n - 2; j >= 0; j-- {
		if maxPrice < s[j] {
			maxPrice = s[j]
		}
		local := maxPrice - s[j]
		if maxProfit < local {
			maxProfit = local
		}
	}
	return maxProfit
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
