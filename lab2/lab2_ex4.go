// Laboratorul 2 - Exercitiul 4
// O goroutine notifica main cand procesarea s-a terminat,
// fara sync.WaitGroup.

package main

import (
	"fmt"
	"time"
)

func processing(done chan<- struct{}) {
	fmt.Println("Procesarea a inceput...")
	time.Sleep(2 * time.Second) // simulam o activitate
	fmt.Println("Procesarea s-a terminat.")
	close(done) // notificam main prin inchiderea canalului
}

func main() {
	done := make(chan struct{})
	go processing(done)

	<-done // main asteapta notificarea
	fmt.Println("Main a primit notificarea de finalizare.")
}
