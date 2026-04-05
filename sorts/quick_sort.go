package sorts

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
		pivotIndex := low + randomGenerator.Intn(high-low+1)
		arr[low], arr[pivotIndex] = arr[pivotIndex], arr[low]

		arr, pivot := partition(arr, low, high)
		// fmt.Println("lvl", "arr", arr,"pivot", pivot, "low",low, "high",high)

		arr = QuickSort(arr, low, pivot)
		arr = QuickSort(arr, pivot+1, high)
	}

	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	// fmt.Println("partition", "arr", arr, "low",low, "high",high)
	pivot := low
	l := pivot + 1
	r := high

	for l <= r {
		if arr[l] < arr[pivot] {
			l++
		} else if arr[r] >= arr[pivot] {
			r--
		} else {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	arr[pivot], arr[r] = arr[r], arr[pivot]
	return arr, r
}
