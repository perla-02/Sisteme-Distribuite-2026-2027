package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct { offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// theta estimeaza diferenta server-client; delta este round-trip-ul
// fara timpul de procesare al serverului.
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func randomNetworkDelay() {
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
}

func measure(cClock, sClock *Clock) (theta, delta time.Duration) {
	t1 := cClock.Now()
	randomNetworkDelay()
	t2 := sClock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := sClock.Now()
	randomNetworkDelay()
	t4 := cClock.Now()
	return estimate(t1, t2, t3, t4)
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 { return -d }
	return d
}

func main() {
	rand.Seed(time.Now().UnixNano())
	serverClock := &Clock{offset: 0}
	initialOffset := 250*time.Millisecond
	clientClock := &Clock{offset: initialOffset}

	fmt.Println("Exercitiul 1 - Algoritmul lui Cristian")
	fmt.Println("Offset initial:", clientClock.offset)
	theta, delta := measure(clientClock, serverClock)
	fmt.Printf("Masurare: theta=%v, delta=%v\\n", theta, delta)

	// theta = server-client, deci il adaugam la offset-ul clientului.
	clientClock.offset += theta
	fmt.Println("Offset dupa corectie:", clientClock.offset)
	fmt.Println("Offset-ul absolut a scazut:",
		absDuration(clientClock.offset) < absDuration(initialOffset))

	var minTheta, minDelta time.Duration
	for i := 0; i < 10; i++ {
		testClient := &Clock{offset: initialOffset} // acelasi offset initial
		th, d := measure(testClient, serverClock)
		if i == 0 || th < minTheta { minTheta = th }
		if i == 0 || d < minDelta { minDelta = d }
	}
	fmt.Println("10 masuratori independente:")
	fmt.Println("Theta minim:", minTheta)
	fmt.Println("Delta minim:", minDelta)
}
