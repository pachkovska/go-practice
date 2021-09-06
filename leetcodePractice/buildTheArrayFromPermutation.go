package leetcodePractice

func buildArray(nums []int) (result []int) {
	for i:=0; i<len(nums); i++ {
		result = append(result, nums[nums[i]])
	}
	return result
}
