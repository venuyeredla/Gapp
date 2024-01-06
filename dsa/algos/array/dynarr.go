package array

import (
	"Gapp/dsa/util"
	"fmt"
)

// Subset sum problem is to find subset of elements that are selected from a given set whose sum adds up to a given number K.
// S,  s={a  a S  }  Sum(s) =K

// 3,2,7,1  = Sum =6
func SubsetSum(A []int, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}
	if sum-A[0] == 0 {
		return true
	}
	if len(A) > 1 {
		including := SubsetSum(A[1:], sum-A[0])
		excluding := SubsetSum(A[1:], sum)

		return including || excluding
	}
	return false
}

func SubsetSumD(A []int, sum int) bool {
	rows := len(A)
	columns := sum + 1
	sm := util.GetMatrix(rows, columns)
	for i := 0; i < len(A); i++ {
		sm[i][0] = 1
	}
	for i := 1; i < columns; i++ {
		if A[0] == i {
			sm[0][i] = 1
		} else {
			sm[0][i] = -1
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if A[i] == j || sm[i-1][j] == 1 || ((j-A[i]) >= 0 && sm[i-1][j-A[i]] == 1) {
				sm[i][j] = 1
			} else {
				sm[i][j] = -1
			}
		}
	}
	if sm[rows-1][columns-1] == 1 {
		return true
	} else {
		return false
	}

}

func LargestSumContiguous(a []int) {
	//StringJoiner stringJoiner=new StringJoiner(", ");
	util.Printable(a, 0, len(a)-1)
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

/*
Input: arr[] = {3, 10, 2, 1, 20} => 3, 10, 20

	Including excluding princieple

Output: 3
*/
func longIncreasingSubseq(A []int) int {

	return 0
}

/*
Input arr[] = {1, 11, 2, 10, 4, 5, 2, 1};
Output: 6 (A Longest Bitonic Subsequence of length 6 is 1, 2, 10, 4, 2, 1)

Increasing then decreasing
*/

func longBitonicSubseq(A []int) int {

	return 0
}
