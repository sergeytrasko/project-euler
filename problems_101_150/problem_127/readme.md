# Project Euler Problem 127 - abc-hits

[https://projecteuler.net/problem=127](https://projecteuler.net/problem=127)

## Solution idea

Before starting the problem, there are some important observations to be made:
- `gcd(a, b) = 1` also implies that `gcd(a, c) = 1` and `gcd(b, c) = 1` (this can be seen from `gcd` definition that `gcd(a+b, b) = gcd(a, b)`)
- if `a`, `b` and `c` are co-prime then they have different prime factors. As result `rad` function is multiplicative

For performance we can pre-compute `rad` for all integers from 1 to 120'000 with an approach similar to primes sieve.
Then just iterate over `a` and `b` and check `rad(a*b*c) < c`. Note that `gcd(a, b) = 1` is moved to the very last step for performance reasons.

## Alternative solutions

## External links
