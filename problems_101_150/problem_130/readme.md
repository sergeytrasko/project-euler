# Project Euler Problem 130 - Composites with prime repunit property

[https://projecteuler.net/problem=130](https://projecteuler.net/problem=130)

## Solution idea

We can take implementation of `A(n)` function from [problem 129 - Repunit divisibility](../problem_129).

Surprisingly, but from now on we can do a brute-force and just for every `n` calculate `k = A(n)` and check if:
- `n` is not prime. This can be done using prime sieve or with probability tests. Golang's `big` package has `ProbablyPrime` function that is based on Baillie–PSW primality test and on Lucas primarity test.
- `(n - 1) % k == 0`

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Repunit
https://en.wikipedia.org/wiki/Primality_test
https://en.wikipedia.org/wiki/Baillie%E2%80%93PSW_primality_test
https://en.wikipedia.org/wiki/Lucas_pseudoprime