package main

import (
	"fmt"
	"sortings/sorts"
)

func main() {
	//arr := []int{5, 2, 3, 1}
	arr := []int{5, 1, 1, 2, 0, 0}
	res := sorts.MergeSort(arr)
	fmt.Println(res)
}
