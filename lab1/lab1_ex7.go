// Lab 1 - Exercitiul 7
// Citeste cuvinte dintr-un fisier si numara aparitiile fiecaruia.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Eroare la deschiderea fisierului:", err)
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Eroare la citirea fisierului:", err)
		return
	}

	fmt.Println("Numarul de aparitii pentru fiecare cuvant:")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
