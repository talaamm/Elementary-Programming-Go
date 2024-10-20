package main

func addPrimeSum(n int) int {
	sum := 0
	for i := n; i >= 1; i-- {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
