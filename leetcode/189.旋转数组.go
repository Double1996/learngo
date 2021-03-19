/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	// 求余数
  newNums := make([]int, len(nums))
    for i, v := range nums {
        newNums[(i+k)%len(nums)] = v
    }
    copy(nums, newNums)
}
// @lc code=end

