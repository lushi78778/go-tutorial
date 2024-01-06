
package main 

import "fmt" 

func main() {

	// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

	// 示例 1:
	// 输入: nums = [-1,0,3,5,9,12], target = 9     
	// 输出: 4       
	// 解释: 9 出现在 nums 中并且下标为 4 

	// 示例 2:
	// 输入: nums = [-1,0,3,5,9,12], target = 2     
	// 输出: -1        
	// 解释: 2 不存在 nums 中因此返回 -1   

	// 提示：
	// 你可以假设 nums 中的所有元素是不重复的。
	// n 将在 [1, 10000]之间。
	// nums 的每个元素都将在 [-9999, 9999]之间。
	var nums []int = []int{-1,0,3,5,9,12}
	target := 9
	fmt.Println(binarySearch3(nums, target))
}


// 二分查找 左闭右闭
func binarySearch(nums []int, target int) int {
    var left int = 0
	var right int = len(nums) - 1

	for left <= right {
		var middle int = (left + right)/ 2
		if(nums[middle] < target){ //中间值小于目标值
			left = middle + 1 // 更新右区间的左边界
		}else if (nums[middle] > target){
			right = middle -1
		}else{
			return middle
		}
	}
    return -1
}

// 二分查找 左闭右开
func binarySearch2(nums []int, target int) int {
    var left int = 0
	var right int = len(nums) //右开
	for left < right {
		var middle int = (left + right)/ 2
		if(nums[middle] < target){ //中间值小于目标值
			left = middle + 1 // 更新右区间的左边界
		}else if (nums[middle] > target){
			right = middle
		}else{
			return middle
		}
	}
    return -1
}

// 二分查找 左开右闭
func binarySearch3(nums []int, target int) int {
    var left int = -1
	var right int = len(nums) -1
	// fmt.Println(right)
	for left < right {
		var middle int = (left + right)/ 2
		if(nums[middle] < target){ //中间值小于目标值
			left = middle // 更新右区间的左边界
		}else if (nums[middle] > target){
			right = middle + 1
		}else{
			return middle
		}
	}
    return -1
}

// 二分查找 左开右开
func binarySearch4(nums []int, target int) int {
    var left int = -1
	var right int = len(nums) 
	// fmt.Println(right)
	for left < right {
		var middle int = (left + right)/ 2
		if(nums[middle] < target){ //中间值小于目标值
			left = middle // 更新右区间的左边界
		}else if (nums[middle] > target){
			right = middle
		}else{
			return middle
		}
	}
    return -1
}
