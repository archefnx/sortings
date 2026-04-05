package sorts

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {

		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
