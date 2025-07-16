//Name: Eduardo Biscaia de Queiroz
//Date: 03/07/2025

//Objectives:
//Create a program that calculates the prime numbers under a given number
//Factorize the number, giving all the prime divisors, prime divisors and
//their repetition or all the divisors of a given number using golang flags

//Examples:
//--all-primes 50
//[2,3,5,7,11,13,17,19,23,29,31,37,41,43,47]

//--factors 50
//[2,5,5]

//--factor-pairs 50
//[(2,1),(5,2)]

//--divisors 50
//[1,2,5,10,50]

package main

import "fmt"

func main() {
	number := 50
	fmt.Println(calcPrimeFactors(number))
}

func checkPrime(number int, primes []int) bool {
	//Check if a number is a prime number
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	}
	for _, prime := range primes {
		if number%prime == 0 {
			return false
		}
	}
	return true
}

func allPrimes(number int) []int {
	listPrimes := []int{}
	for i := 1; number >= i; i++ {
		if checkPrime(i, listPrimes) {
			listPrimes = append(listPrimes, i)
		}
	}
	return listPrimes
}

func calcPrimeFactors(number int) []int {
	divisor := 1
	primes := []int{}
	primeFactors := []int{}
	for number > 1 {
		for !checkPrime(divisor, primes) {
			divisor++
		}
		for number%divisor == 0 {
			primeFactors = append(primeFactors, divisor)
			number /= divisor
		}
		divisor++
	}
	return primeFactors
}
