package sorts

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	arr = MSort(left, right)

	return arr
}

func MSort(left, right []int) []int {
	var l, r, resIdx int
	size := len(left) + len(right)
	res := make([]int, size)

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res[resIdx] = left[l]
			l++
		} else {
			res[resIdx] = right[r]
			r++
		}
		resIdx++
	}

	for l < len(left) {
		res[resIdx] = left[l]
		resIdx++
		l++
	}

	for r < len(right) {
		res[resIdx] = right[r]
		resIdx++
		r++
	}

	return res
}
