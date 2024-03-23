package main

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
