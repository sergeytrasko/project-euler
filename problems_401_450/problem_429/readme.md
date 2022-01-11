# Project Euler Problem 429 - Sum of squares of unitary divisors

[https://projecteuler.net/problem=429](https://projecteuler.net/problem=429)

## Solution idea

For this problem we need to solve 2 sub-problems:
- Factor `100'000'000!` into prime factors
- Sum unitary divisors based on prime factorization

Factorization is quite easy - we need to generate a sieve of primes below `100'000'000` and notice that for `p^e` to divide `n` we should have `e = [n/e] + [n/e^2] + [n/e^3] + ...`. At some point of time `e^k > n` and the sum converges.

So in `O(n*logn)` we can have prime factorization of `n!`.

Next observation: `d` is an unitary divisor of `n` if `d = p^e` and `p^e` is a maximal divisor of `p` in `n` factorization.
This means that we need to deal only with individual prime factors of `n!`.

In turns out that function to sum unitary divisors is multiplicative:
- consider `b(n,k)` as sum of `k-th` powers of unitary divisors
- then if `n = p*q` and `gcd(p, q) = 1` then `b(n, k) = b(p*q, k) = b(p, k) * b(q, k)`
- futher more: `b(p^e, k) = p^(k*e) + 1` where `p` is prime

This gives us all building blocks for the solution.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Divisor_function

https://mathworld.wolfram.com/UnitaryDivisorFunction.html

https://oeis.org/A034676