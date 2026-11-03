// Laboratorul 2 - Exercitiul 7
// Doi roboti isi transmit periodic coordonatele 2D.
// La terminarea traseului, fiecare trimite un mesaj de oprire.
// Fiecare robot citeste mesajele pana afla ca si celalalt s-a oprit,
// apoi raporteaza terminarea catre observer (main).

package main

import (
	"fmt"
	"time"
)

type Message struct {
	x, y   int
	stopped bool
}

type Report struct {
	robot string
}

func robot(
	name string,
	route [][2]int,
	out chan<- Message,
	in <-chan Message,
	reports chan<- Report,
) {
	// Trimitem toate pozitiile traseului. Canalul este buffered pentru
	// a permite celor doi roboti sa transmita in acelasi timp.
	for _, point := range route {
		time.Sleep(200 * time.Millisecond)
		out <- Message{x: point[0], y: point[1]}
		fmt.Printf("%s a trimis pozitia (%d, %d)\n",
			name, point[0], point[1])
	}

	// Mesajul de oprire este trimis dupa toate pozitiile proprii.
	out <- Message{stopped: true}
	fmt.Println(name, "si-a terminat traseul si a notificat celalalt robot.")

	// Canalul pastreaza ordinea: toate pozitiile trimise anterior
	// vor fi citite inaintea mesajului stopped.
	peerStopped := false
	for !peerStopped {
		msg := <-in
		if msg.stopped {
			peerStopped = true
			fmt.Println(name, "a primit notificarea ca celalalt robot s-a oprit.")
		} else {
			fmt.Printf("%s a primit pozitia celuilalt robot: (%d, %d)\n",
				name, msg.x, msg.y)
		}
	}

	reports <- Report{robot: name}
}

func main() {
	route1 := [][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}}
	route2 := [][2]int{{0, 0}, {1, 1}, {2, 2}}

	// Capacitatea acopera toate pozitiile si mesajul de oprire.
	ch12 := make(chan Message, len(route1)+1)
	ch21 := make(chan Message, len(route2)+1)
	reports := make(chan Report, 2)

	go robot("Robot 1", route1, ch12, ch21, reports)
	go robot("Robot 2", route2, ch21, ch12, reports)

	// Observer asteapta rapoartele ambilor roboti.
	first := <-reports
	second := <-reports
	fmt.Println("Observer: ambii roboti au terminat:", first.robot, "si", second.robot)
}
