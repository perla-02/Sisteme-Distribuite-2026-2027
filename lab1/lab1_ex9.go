// Lab 1 - Exercitiul 9
// Gaseste in O(n) indicii a doua elemente a caror suma este s.
// Se foloseste un map pentru a retine valorile deja parcurse.

package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 8, 4, 9}
	target := 15

	seen := make(map[int]int)

	fmt.Println("Array:", numbers)
	fmt.Println("Suma cautata:", target)
	fmt.Println("Perechi de indici:")

	found := false
	for i, value := range numbers {
		needed := target - value

		if previousIndex, ok := seen[needed]; ok {
			fmt.Printf("%d %d\n", previousIndex, i)
			found = true
		}

		seen[value] = i
	}

	if !found {
		fmt.Println("Nu exista o pereche care sa aiba suma ceruta.")
	}
}
