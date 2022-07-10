package main

import "fmt"

func main() {

	result := twoSum([]int{1,2,3,4,5}, 7);
	fmt.Println(result);

	result1 := twoSum([]int{1,2,3,4,5}, 8);
	fmt.Println(result1);

}

func twoSum(nums []int, target int) []int {
	
	hashMap := make(map[int]int);

	for i := 0; i < len(nums) - 1; i++ {
		if _, prs := hashMap[nums[i]]; prs {
			return []int{ hashMap[nums[i]], i};
		}

		hashMap[target - nums[i]] = i;
	} 

	return []int{};
}