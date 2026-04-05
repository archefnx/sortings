package sorts

func CountingSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	count := make([]int, max-min+1)
	for _, num := range nums {
		count[num-min]++
	}

	index := 0
	for i := 0; i < len(count); i++ {
		for count[i] > 0 {
			nums[index] = i + min
			index++
			count[i]--
		}
	}

	return nums
}
