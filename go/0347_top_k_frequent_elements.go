package main

import (
	"fmt"
	"reflect"
)

func topKFrequent(nums []int, k int) (res []int) {
	countMap := map[int]int{}
	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
		}
	}

	bucket := make([][]int, len(nums)+1)
	for num, count := range countMap {
		bucket[count] = append(bucket[count], num)
	}

	for i := len(bucket) - 1; i > 0; i-- {
		res = append(res, bucket[i]...)
		if len(res) == k {
			return
		}
	}

	return
}

func main() {
	nums := []int{1, 1, 2, 2, 3}
	k := 2
	expectedResult := []int{1, 2}
	result := topKFrequent(nums, k)

	if reflect.DeepEqual(result, expectedResult) {
		fmt.Println("The results are equal.")
	} else {
		fmt.Println("The results are not equal.")
	}
}
