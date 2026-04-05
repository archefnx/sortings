package main

import (
	"fmt"
	"sortings/sorts"
)

func main() {
	//arr := []int{5, 2, 3, 1}
	arr := []int{5, 1, 1, 2, 0, 0}
	//res := sorts.MergeSort(arr)
	//res := sorts.QuickSort(arr, 0, len(arr)-1)
	//res := sorts.InsertionSort(arr)
	//res := sorts.SelectionSort(arr)
	res := sorts.CountingSort(arr)
	fmt.Println(res)
}
