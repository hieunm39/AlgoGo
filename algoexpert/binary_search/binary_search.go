package binarysearch

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.


func Search(nums []int, target int) int {


	left := 0;
	right := len(nums) -1;

	for left <= right {
		
		mid := left + (right - left) / 2;

		if nums[mid] == target {
			return mid;
		}

		if nums[mid] < target {
			left = mid + 1;
		}

		if nums[mid] > target {
			right = mid - 1;
		}
	}
	
	return - 1;
}