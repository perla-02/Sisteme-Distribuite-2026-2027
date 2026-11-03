package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct { counter int }

func (l *LamportClock) Tick() { l.counter++ }
func (l *LamportClock) Send() { l.counter++ }
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter { l.counter = ts }
	l.counter++
}

type Message struct { sender, ts int }
type Event struct {
	process int
	internal_counter int // timestamp Lamport
	kind string // "internal", "send", "recv"
}

// Ordoneaza crescator dupa timestamp; egalitatile sunt departajate determinist.
func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	if a.process != b.process { return a.process < b.process }
	return a.kind < b.kind
}

func process(id int, in <-chan Message, out chan<- Message,
	events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{}

	time.Sleep(time.Duration(rand.Intn(8)+1)*time.Millisecond)
	clock.Tick()
	events <- Event{process:id, internal_counter:clock.counter, kind:"internal"}

	time.Sleep(time.Duration(rand.Intn(8)+1)*time.Millisecond)
	clock.Send()
	events <- Event{process:id, internal_counter:clock.counter, kind:"send"}
	out <- Message{sender:id, ts:clock.counter}

	msg := <-in
	clock.Receive(msg.ts)
	events <- Event{process:id, internal_counter:clock.counter, kind:"recv"}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch12 := make(chan Message)
	ch23 := make(chan Message)
	ch31 := make(chan Message)
	events := make(chan Event, 9)
	var wg sync.WaitGroup
	wg.Add(3)

	go process(1, ch31, ch12, events, &wg)
	go process(2, ch12, ch23, events, &wg)
	go process(3, ch23, ch31, events, &wg)

	wg.Wait()
	close(events)
	var allEvents []Event
	for event := range events { allEvents = append(allEvents, event) }
	sort.Slice(allEvents, func(i, j int) bool { return less(allEvents[i], allEvents[j]) })

	fmt.Println("Exercitiul 2 - Evenimente Lamport ordonate crescator")
	for _, e := range allEvents {
		fmt.Printf("P%d | timestamp=%d | %s\\n", e.process, e.internal_counter, e.kind)
	}
}
