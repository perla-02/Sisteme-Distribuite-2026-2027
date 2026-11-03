// Lab 1 - Exercitiul 3
// Afiseaza un slice cu numerele pare dintr-un sir.

package main

import "fmt"

func main() {
	numbers := []int{12, 7, 4, 9, 16, 3, 20, 11, 8}
	evens := []int{}

	for _, number := range numbers {
		if number%2 == 0 {
			evens = append(evens, number)
		}
	}

	fmt.Println("Sirul initial:", numbers)
	fmt.Println("Numere pare:", evens)
}
