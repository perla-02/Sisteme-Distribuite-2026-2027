// Laboratorul 2 - Exercitiul 3
// main primeste doua mesaje pe acelasi canal.

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Primul mesaj"
		ch <- "Al doilea mesaj"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("Main a primit:", <-ch)
	}
}
