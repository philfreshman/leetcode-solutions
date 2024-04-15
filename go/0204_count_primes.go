package main

// time: O(n log log n)
// space: O(n)

// countPrimes implements the Sieve of Eratosthenes algorithm
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
