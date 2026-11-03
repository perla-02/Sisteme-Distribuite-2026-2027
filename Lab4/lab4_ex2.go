// Laboratorul 4 - Exercitiul 2
// Generalized Flooding pentru un graf tare conex.
//
// edges[i] contine indicii nodurilor catre care nodul i trimite mesaje.
// ID-urile sunt valori unice; rezultatul este ID-ul maxim.

package main

import "fmt"

// GeneralizedFlooding propaga maximul ID-urilor pe muchiile grafului.
// Graful trebuie sa fie tare conex.
func GeneralizedFlooding(ids []int, edges [][]int) int {
	n := len(ids)
	if n == 0 || len(edges) != n {
		return -1
	}

	// Verificam unicitatea ID-urilor si validitatea indicilor.
	seen := make(map[int]bool)
	for _, id := range ids {
		if seen[id] {
			return -1
		}
		seen[id] = true
	}
	for from := range edges {
		for _, to := range edges[from] {
			if to < 0 || to >= n {
				return -1
			}
		}
	}

	// maxID[i] reprezinta cea mai mare valoare cunoscuta de nodul i.
	maxID := make([]int, n)
	copy(maxID, ids)

	// Runde sincrone: toate actualizarile folosesc valorile din runda curenta.
	for {
		next := make([]int, n)
		copy(next, maxID)
		changed := false

		for from := 0; from < n; from++ {
			for _, to := range edges[from] {
				if maxID[from] > next[to] {
					next[to] = maxID[from]
					changed = true
				}
			}
		}

		maxID = next
		if !changed {
			break
		}
	}

	// Intr-un graf tare conex, toate nodurile ajung sa cunoasca maximul global.
	return maxID[0]
}

func main() {
	ids := []int{5, 12, 8, 20}

	// Graful orientat: fiecare nod are muchii catre vecinii indicati.
	// Acest exemplu este tare conex.
	edges := [][]int{
		{1, 2}, // nodul 0 trimite la 1 si 2
		{2, 3}, // nodul 1 trimite la 2 si 3
		{3},    // nodul 2 trimite la 3
		{0},    // nodul 3 trimite la 0
	}

	leader := GeneralizedFlooding(ids, edges)
	fmt.Println("ID-urile nodurilor:", ids)
	fmt.Println("Liderul ales:", leader)
}
