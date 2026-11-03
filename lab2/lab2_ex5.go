// Laboratorul 2 - Exercitiul 5
// Un goroutine send trimite un mesaj unui goroutine receive.
// Receive returneaza mesajul, iar main il afiseaza.

package main

import "fmt"

func send(ch chan<- string) {
	ch <- "salut"
}

func receive(ch <-chan string) string {
	return <-ch
}

func main() {
	ch := make(chan string)
	go send(ch)

	message := receive(ch)
	fmt.Println("Mesaj primit:", message)
}
