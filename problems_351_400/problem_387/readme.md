# Project Euler Problem 387 - Harshad Numbers

[https://projecteuler.net/problem=387](https://projecteuler.net/problem=387)

## Solution idea

Instead of checking each and every number for prime and strong right truncateable Harshad numebrs we can do it other way around.
Let's generate all these numbers, by starting with 1, 2, 3, ... 9 and appending digits to the right applying rules for right truncateable Harshad numebrs. On every recursion step we check the conditions for prime number.

Next problem is how to check if a number is prime. Of course we can get a prime number sieve up to 10^7 and then check every of them.
However here we can use probability test that gives 100% result for numbers below 2^64
Golang's `big` package has `ProbablyPrime` function that is based on Baillie–PSW primality test and on Lucas primarity test - this saves us a lot of time and effort.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Harshad_number
https://en.wikipedia.org/wiki/Primality_test
https://en.wikipedia.org/wiki/Baillie%E2%80%93PSW_primality_test
https://en.wikipedia.org/wiki/Lucas_pseudoprime