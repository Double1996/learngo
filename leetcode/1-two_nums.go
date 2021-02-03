package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{i, v}
		}
		m[n] = i
	}
	return nil
}
