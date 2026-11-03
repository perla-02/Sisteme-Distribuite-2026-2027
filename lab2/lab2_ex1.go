// Laboratorul 2 - Exercitiul 1
// O goroutine diferita de main trimite mesajul "salut" catre main prin channel.

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "salut"
	}()

	message := <-ch
	fmt.Println("Mesaj primit in main:", message)
}
