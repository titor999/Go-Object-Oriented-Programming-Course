package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Input the diapason numbers >>>")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Following are the prime numbers ")
	fmt.Printf("smaller than or equal to %v\n", n)
	sieveOfEratosthenes(n)
}

func sieveOfEratosthenes(n int) {
	// Create a boolean array "prime[0…n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if I am Not a prime, else true.
	prime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If prime[p] is not changed, then it is a prime
		if prime[p] {
			// Update all multiples of p
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}

		// Print all prime numbers
		for i := 2; i <= n; i++ {
			if prime[i] {
				fmt.Printf("%v ", i)
			}
		}
	}
}
