package main

import (
	"first-prj/sorter"
	"fmt"
	"math/rand"
)

func main() {
	sorter.Help()

	arr := rand.Perm(100)
	// arr := []int{1, 7, 6, 2, 3, 0, 5, 4, 9, 8}
	fmt.Println(arr)

	// sortedArr := sorter.Partition(arr, 0, len(arr)-1, len(arr)-1)
	sortedArr := sorter.QuickSort(arr)
	fmt.Println(sortedArr)
	fmt.Println(arr)
}
