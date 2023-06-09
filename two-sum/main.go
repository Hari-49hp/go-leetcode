package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

func twoSum(nums []int, target int) []int {
	length := len(nums)

	first := length - 1
	second := length - 2

	for first >= 0 && second >= 0 {
		res := nums[first] + nums[second]
		if target == res {
			result := []int{first, second}
			sort.Ints(result)
			return result
		}

		if second == 0 {
			first--
			second = first - 1
		} else {
			second--
		}

		length = length - 1
	}

	return nil
}
