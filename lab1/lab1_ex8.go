// Lab 1 - Exercitiul 8
// Grupeaza angajatii dupa id-ul departamentului.

package main

import "fmt"

func main() {
	employees := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	departments := make(map[int][]string)

	for name, departmentID := range employees {
		departments[departmentID] = append(departments[departmentID], name)
	}

	for departmentID, names := range departments {
		fmt.Printf("Departamentul %d: %v\n", departmentID, names)
	}
}
