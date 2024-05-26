package twosum

func TwoSum(nums []int, target int) []int {
	var out []int
	var search int
	for i, v := range nums {
		search = target - v
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == search {
				return []int{i, j}
			}
		}
	}
	return out
}
