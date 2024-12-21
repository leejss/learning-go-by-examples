package main

import "fmt"

func double(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {

	// 1. make a slice using make or slice literal
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums1, len(nums1), cap(nums1)) // [1 2 3 4 5] 5 5
	// make(<type of slice>, <length>, <capacity>)
	nums2 := make([]int, 5)

	// 2. append to a slice
	// Go will allocate a new underlying array (usually doubling capacity) and copy the existing elements over.
	nums1 = append(nums1, 6)                   // capacity is doubled
	fmt.Println(nums1, len(nums1), cap(nums1)) // [1 2 3 4 5 6] 6 10
	// capacity가 증가하게 되면, 기존 배열에 대한 reference는 끊어진다.

	// 두 slice 이어붙이기
	nums3 := append(nums1, nums2...)
	fmt.Println(nums3, len(nums3), cap(nums3)) // [1 2 3 4 5 6 0 0 0 0] 10 10

	// 3. slice는 reference type이다.
	double(nums1)
	fmt.Println(nums1) // [2 4 6 8 10 12]

	// 4. nil slice and empty slice
	var nums4 []int
	fmt.Println(nums4, len(nums4), cap(nums4)) // [] 0 0

	nums5 := []int{}
	fmt.Println(nums5, len(nums5), cap(nums5)) // [] 0 0

	// check nil
	fmt.Println(nums4 == nil) // true
	fmt.Println(nums5 == nil) // false

	// 5. copy slice without reference
	nums6 := []int{1, 2, 3, 4, 5}
	nums7 := make([]int, len(nums6))
	copy(nums7, nums6)
	fmt.Println(nums7) // [1 2 3 4 5]

	// 6. removing or inserting at index
	nums8 := []int{1, 2, 3, 4, 5}
	removingIndex := 2
	nums8 = append(nums8[:removingIndex], nums8[removingIndex+1:]...)
	fmt.Println(nums8) // [1 2 4 5]

}
