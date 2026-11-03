// Lab 1 - Exercitiul 2
// Sorteaza crescator un array folosind Bubble Sort.

package main

import "fmt"

func main() {
	arr := [8]int{7, 2, 9, 1, 5, 3, 8, 4}

	fmt.Println("Array initial:", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Array sortat crescator:", arr)
}
