package array

import "fmt"

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

func linearSearch(input []int, key int) bool {
	var contains bool = false
	for i := 0; i < len(input); i++ {
		if input[i] == key {
			contains = true
			break
		}
	}
	return contains
}

func linearSearchR(input []int, lastIdx int, key int) int {
	if lastIdx == -1 {
		return -1
	} else if input[lastIdx] == key {
		return lastIdx
	} else {
		return linearSearchR(input, lastIdx-1, key)
	}
}

func binarySearch(input []int, key int) bool {
	l := 0
	r := len(input)
	for l < r {
		m := l + ((r - l) / 2)
		if key > input[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	if input[l] == key {
		return true
	} else {
		return false
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
			binarySearchR(input, m+1, right, key)
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
			if input[minIdx] != input[i] {
				Swap(input, minIdx, i)
			}
		}
		break
	case Insertion:
		for i := 1; i < len(input); i++ {
			j := i - 1
			for j >= 0 && input[i] < input[j] {
				j--
			}
			j++
			if j != i {
				temp := input[i]
				for k := i - 1; k > j; k-- {
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
