// Laboratorul 4 - Exercitiul 1
// Flooding pentru un inel unidirectional.
// Fiecare nod isi trimite propriul ID. Un ID primit este retransmis
// numai daca este mai mare decat ID-ul local. ID-ul maxim se intoarce
// la nodul de la care a pornit, iar acesta anunta alegerea liderului.

package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

// Realizeaza operatiile interne ale unui nod din inel.
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// Fiecare nod isi introduce propriul ID in inel.
	out <- Msg{Value: id}

	for {
		msg := <-in

		if msg.Done {
			// Anuntul de alegere este transmis mai departe o singura data.
			out <- msg
			return
		}

		switch {
		case msg.Value == id:
			// ID-ul propriu a parcurs tot inelul: acest nod este lider.
			*isLeader = true
			out <- Msg{Value: id, Done: true}
			return
		case msg.Value > id:
			// Candidatul are un ID mai mare; il retransmitem.
			out <- msg
		default:
			// Candidatul este mai mic: il ignoram.
		}
	}
}

// RingFlooding returneaza ID-ul liderului (ID-ul maxim).
func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	// Algoritmul presupune ID-uri unice.
	seen := make(map[int]bool)
	for _, id := range ids {
		if seen[id] {
			return -1
		}
		seen[id] = true
	}

	channels := make([]chan Msg, n)
	for i := range channels {
		// Bufferul permite fiecarui nod sa-si trimita initial ID-ul.
		channels[i] = make(chan Msg, n+2)
	}

	isLeader := make([]bool, n)
	var wg sync.WaitGroup

	for i, id := range ids {
		wg.Add(1)
		in := channels[(i-1+n)%n]
		out := channels[i]
		go node(id, in, out, &isLeader[i], &wg)
	}

	wg.Wait()

	for i, elected := range isLeader {
		if elected {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{5, 12, 8, 20, 3}
	leader := RingFlooding(ids)
	fmt.Println("ID-urile nodurilor:", ids)
	fmt.Println("Liderul ales:", leader)
}
