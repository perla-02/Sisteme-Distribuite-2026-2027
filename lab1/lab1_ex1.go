// Lab 1 - Exercitiul 1
// Citeste un numar natural si verifica daca este prim.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Introduceti un numar natural: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println(n, "nu este numar prim.")
		return
	}

	isPrime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(n, "este numar prim.")
	} else {
		fmt.Println(n, "nu este numar prim.")
	}
}
