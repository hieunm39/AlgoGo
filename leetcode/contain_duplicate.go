package main

import "fmt"

func main() {
	nums := []int{1,2,3,1};
	nums2 := []int{1,2,3,4};

	fmt.Println(containsDuplicate(nums)); // true
	fmt.Println(containsDuplicate(nums2)); // false
}

func containsDuplicate(nums []int) bool {
    
    hashMap := make(map[int]bool)
    
    for _, v := range nums {
        
        if hashMap[v] == true {
            
            return true;
        }
        hashMap[v] = true;
        
    }
    return false
}