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

import (
	"fmt"
	"math"
	"sort"
)

type divisor struct {
	factor int
	power  int
}

func main() {
	number := 50
	fmt.Println(allPowers(factorsPowers(calcPrimeFactors(number))))
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

func factorsPowers(number int) []divisor {
	prev_factor := 1
	factors := calcPrimeFactors(number)
	factorsAndPowers := []divisor{}
	div := divisor{}
	for _, cur_factor := range factors {
		if cur_factor != prev_factor {
			div = divisor{factor: cur_factor, power: 1}
			factorsAndPowers = append(factorsAndPowers, div)
			prev_factor = cur_factor
		} else {
			factorsAndPowers[len(factorsAndPowers)-1].power++
		}
	}
	return factorsAndPowers
}

func allPowers(factorsAndPowers []divisor) [][]int {
	powers := [][]int{}
	for _, primeFactor := range factorsAndPowers {
		factorPowers := []int{}
		for i := 0; i <= primeFactor.power; i++ {
			factorPowers = append(factorPowers, int(math.Pow(float64(primeFactor.factor), float64(i))))
		}
		powers = append(powers, factorPowers)
	}
	return powers
}

func allDivisors(powers [][]int) []int {
	listOfDivisors := []int{1}
	for _, power := range powers {
		newListOfDivisors := []int{}
		for _, number := range power {
			for _, div := range listOfDivisors {
				newListOfDivisors = append(newListOfDivisors, number*div)
			}
		}
		listOfDivisors = newListOfDivisors
	}
	sort.Ints(listOfDivisors)
	return listOfDivisors
}
