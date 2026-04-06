package sorts

import (
	"math"
	"slices"
)

func BucketSort(nums []int) []int {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	if min == max {
		return nums
	}

	n := len(nums)

	bucketSize := int(math.Ceil(float64(max-min) / float64(n-1)))
	bucketCount := (max-min)/bucketSize + 1
	buckets := make([][]int, bucketCount)

	for _, num := range nums {
		idx := (num - min) / bucketSize
		buckets[idx] = append(buckets[idx], num)
	}

	ans := make([]int, 0, len(nums))

	for _, b := range buckets {
		if len(b) > 0 {
			slices.Sort(b)
			ans = append(ans, b...)
		}
	}

	return ans
}
