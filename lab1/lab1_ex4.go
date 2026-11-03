// Lab 1 - Exercitiul 4
// Elimina elementul de la indexul citit, pastrand ordinea celorlalte elemente.

package main

import "fmt"

func main() {
	values := []float32{1.5, 2.5, 3.5, 4.5, 5.5}

	fmt.Println("Slice initial:", values)
	fmt.Printf("Introduceti indexul de eliminat (0-%d): ", len(values)-1)

	var index int
	fmt.Scan(&index)

	if index < 0 || index >= len(values) {
		fmt.Println("Index invalid.")
		return
	}

	values = append(values[:index], values[index+1:]...)
	fmt.Println("Slice dupa eliminare:", values)
}
