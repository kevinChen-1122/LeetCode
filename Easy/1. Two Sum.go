package Easy

// https://leetcode.com/problems/two-sum/description/

func BruteForce(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func Map(nums []int, target int) []int {
	slice := make(map[int]int)
	for index, num := range nums {
		if previousIndex, ok := slice[target-num]; ok {
			return []int{previousIndex, index}
		}
		slice[num] = index
	}
	return []int{}
}
