package main

import "fmt"

// slice append

func main() {

	slic := make([]int, 5)
	// fmt.Printf("%p\n", slic)

	for i := 0; i < 10; i++ {
		slic = append(slic, i)
		// fmt.Printf("%2d : slic ptr -> %p \n", i, slic)
	}

	output(slic, "slic")
	rotate(slic, 5)
}

func rotate(nums []int, k int) {
	output(nums, "nums")

	n := len(nums)
	k = k % n
	num1 := nums[:n-k]
	num2 := nums[n-k:]
	output(num2, "num2")
	output(num1, "num1")

	// fmt.Println(num1, num2)
	nums = append(num2, num1...)
	output(nums, "nums2nd")

}

func output(nums []int, name string) {
	fmt.Printf("%s array ptr: %p-> %p , len(%d), cap(%d) \n", name, nums, &nums, len(nums), cap(nums))
}
