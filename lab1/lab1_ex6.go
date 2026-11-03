// Lab 1 - Exercitiul 6
// Afiseaza numele in ordinea descrescatoare a inaltimii.

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
}

func main() {
	peopleMap := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	people := make([]Person, 0, len(peopleMap))
	for name, height := range peopleMap {
		people = append(people, Person{name: name, height: height})
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})

	fmt.Println("Persoane in ordinea descrescatoare a inaltimii:")
	for _, person := range people {
		fmt.Printf("%s: %d cm\n", person.name, person.height)
	}
}
