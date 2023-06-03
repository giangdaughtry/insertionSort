package main

import "fmt"

func main() {
	arr := [...]int{3, 54, 64, 13, 6, 876, 24, 13, 5, 978}
	sizeArray := len(arr)
	insertionSort(arr[:], sizeArray)
	fmt.Print(arr)

}

func insertionSort(arr []int, size int) {
	for i := 1; i < size; i++ {
		max := arr[i]
		j := i - 1
		for (j >= 0) && (max < arr[j]) {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = max
	}
}
