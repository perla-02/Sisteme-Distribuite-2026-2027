// Laboratorul 4 - Exercitiul 3
// Flooding de medie pentru vot majoritar.
//
// Formula folosita conform suportului de laborator:
// x_i(t+1) = (x_i(t) + suma valorilor vecinilor) / (grad(i) + 1)
//
// Pentru ca rezultatul consensului sa fie media aritmetica initiala,
// matricea ponderilor trebuie sa fie dublu stocastica. Exemplul de mai jos
// foloseste un inel neorientat regulat, unde fiecare nod are gradul 2.

package main

import (
	"fmt"
	"math"
)

// FloodingAverage intoarce media finala estimata, valorile nodurilor,
// numarul de iteratii si daca s-a atins criteriul de oprire.
func FloodingAverage(initial []float64, neighbors [][]int, epsilon float64, maxIterations int) (float64, []float64, int, bool) {
	n := len(initial)
	if n == 0 || len(neighbors) != n || epsilon <= 0 || maxIterations <= 0 {
		return 0, nil, 0, false
	}

	current := make([]float64, n)
	copy(current, initial)

	// Verificam indicii vecinilor.
	for i := range neighbors {
		for _, neighbor := range neighbors[i] {
			if neighbor < 0 || neighbor >= n {
				return 0, nil, 0, false
			}
		}
	}

	for iteration := 1; iteration <= maxIterations; iteration++ {
		next := make([]float64, n)

		// Actualizarea sincrona: folosim numai valorile din current.
		for i := 0; i < n; i++ {
			sum := current[i]
			for _, neighbor := range neighbors[i] {
				sum += current[neighbor]
			}
			next[i] = sum / float64(len(neighbors[i])+1)
		}

		current = next

		// Criteriul global de oprire din suport: max(x)-min(x) < epsilon.
		minValue, maxValue := current[0], current[0]
		for _, value := range current[1:] {
			if value < minValue {
				minValue = value
			}
			if value > maxValue {
				maxValue = value
			}
		}

		if maxValue-minValue < epsilon {
			consensus := 0.0
			for _, value := range current {
				consensus += value
			}
			consensus /= float64(n)
			return consensus, current, iteration, true
		}
	}

	// Daca nu s-a atins toleranta, intoarcem media valorilor curente,
	// dar marcam faptul ca algoritmul nu a convergat in limita aleasa.
	consensus := 0.0
	for _, value := range current {
		consensus += value
	}
	consensus /= float64(n)

	return consensus, current, maxIterations, false
}

func main() {
	// Patru voturi: 1 = DA, 0 = NU.
	initial := []float64{1, 0, 1, 1}

	// Inel neorientat cu 4 noduri; fiecare nod are doi vecini.
	neighbors := [][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}

	consensus, values, iterations, converged := FloodingAverage(
		initial,
		neighbors,
		1e-9,
		10000,
	)

	fmt.Printf("Valorile finale: %.6f\n", values)
	fmt.Printf("Consens estimat: %.6f\n", consensus)
	fmt.Println("Iteratii:", iterations)
	fmt.Println("A convergat:", converged)

	if math.Abs(consensus-0.5) < 1e-9 {
		fmt.Println("Rezultat: egalitate")
	} else if consensus > 0.5 {
		fmt.Println("Rezultat vot: DA (majoritate)")
	} else {
		fmt.Println("Rezultat vot: NU (majoritate)")
	}
}
