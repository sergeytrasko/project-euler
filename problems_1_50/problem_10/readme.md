# Project Euler Problem 10 - Summation of primes

[https://projecteuler.net/problem=10](https://projecteuler.net/problem=10)

## Solution idea

Again brute-force solution - iterate from 2 to 2'000'000, check if number is prime and if yes - add it to the sum

## Alternative solutions

Use sieve algorithms for generating big sets of prime numbers, e.g. sieve of Eratosthenes.
Refer to [utils.go](../../utils/utils.go) for the implementation details.

## External links

https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes