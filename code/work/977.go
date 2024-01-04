
package main 

import "fmt" 
import "sort" 

func main() {

	// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

	// 示例 1：
	// 输入：nums = [-4,-1,0,3,10]
	// 输出：[0,1,9,16,100]
	// 解释：平方后，数组变为 [16,1,0,9,100]
	// 排序后，数组变为 [0,1,9,16,100]

	// 示例 2：
	// 输入：nums = [-7,-3,2,3,11]
	// 输出：[4,9,9,49,121]
	 
	
	// 提示：
	// 1 <= nums.length <= 104
	// -104 <= nums[i] <= 104
	// nums 已按 非递减顺序 排序
	 
	
	// 进阶：
	// 请你设计时间复杂度为 O(n) 的算法解决本问题

	var nums []int = []int{-7,-3,2,3,11}
	fmt.Println(sortedSquares2(nums))
}

// 暴力排序
func sortedSquares(nums []int) []int {
	var res []int
	for i:=0; i < len(nums); i++ {
		res = append(res, nums[i] * nums[i])
	}
	sort.Ints(res)
	return res
}


// 双指针排序
func sortedSquares2(nums []int) []int {
	var res []int
	for i:=0; i < len(nums); i++ {
		res = append(res, nums[i] * nums[i])
	}
	fmt.Println(res)

	var res2 []int
	left := 0
	for res[left] > res[left+1]{
		left++	
	}
	right := left + 1
	for left >= 0 && right < len(res) {
		if res[left] > res[right]{
			res2 = append(res2,res[right])
			if right < len(res)-1{
				right++
			}
		}else{
			res2 = append(res2,res[left])
			if left > 0{
				left--
			}
		}
	}

	return res2
}
