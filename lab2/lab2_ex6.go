// Laboratorul 2 - Exercitiul 6
// Doi roboti pornesc simultan si alearga spre finish.
// select identifica primul mesaj de sosire.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	robot string
	time  time.Duration
}

func robot(name string, results chan<- Result, start <-chan struct{}) {
	<-start // ambii roboti asteapta semnalul comun de start

	duration := time.Duration(500+rand.Intn(1500)) * time.Millisecond
	time.Sleep(duration)

	results <- Result{robot: name, time: duration}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	results := make(chan Result, 2)
	start := make(chan struct{})

	go robot("Robot 1", results, start)
	go robot("Robot 2", results, start)

	fmt.Println("Observer: start!")
	close(start) // porneste ambii roboti

	// Primul rezultat primit prin canal reprezinta castigatorul.
	select {
	case winner := <-results:
		fmt.Printf("%s a ajuns primul la finish (timp: %v)!\n",
			winner.robot, winner.time)
	}
}
