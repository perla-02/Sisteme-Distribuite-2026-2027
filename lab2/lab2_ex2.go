// Laboratorul 2 - Exercitiul 2
// Goroutine main trimite mesajul "salut" catre o alta goroutine.

package main

import "fmt"

func worker(ch <-chan string) {
	message := <-ch
	fmt.Println("Mesaj primit in goroutine:", message)
}

func main() {
	ch := make(chan string)
	go worker(ch)

	ch <- "salut"
}
