// Lab 1 - Exercitiul 5
// Inverseaza slice-ul in-place, fara a crea un alt slice.

package main

import "fmt"

func main() {
	words := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	fmt.Println("Slice initial:", words)

	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}

	fmt.Println("Slice inversat:", words)
}
