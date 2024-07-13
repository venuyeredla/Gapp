package dynamic

import "Gapp/dsa/utils"

func Rob(nums []int, idx int) int {
	if idx < 0 {
		return 0
	}
	includeSum := nums[idx] + Rob(nums, idx-2)
	excludeSum := Rob(nums, idx-1)

	return utils.MaxOf(includeSum, excludeSum)
}
