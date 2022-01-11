# Project Euler Problem 549 - Divisibility of factorials

[https://projecteuler.net/problem=549](https://projecteuler.net/problem=549)

## Solution idea

Having the task to compute `S(10^8)` means that it is not possible to calculate each individual `s(n)` in brute force way as in worst case `s(n) = n` for primes.

My approach is:
- find factorization of all numbers from `2` to `10^8`. This can be done using "smallest prime factor" approach (refer to [problem 159 - Digital root sums of factorisations](../../problems_151_200/problem_159))
- having factorization of `n` in place we can construct a factorial `y!` that is divisible by `n`

Let's look in example how to do it:
- take `n = 2^4 * 3^3 = 432` that has 2 prime factors - `2` and `3`
- find all numbers divisible by `2` that exponents for factor of `2` sum up to `4` - these are `2`, `2^2=4` and `2^3=8`. So we need to have at least `y > =8` to have `2^4 | y!`
- similarly - take all numbers divisible by `3` - `3` and `3^2 = 9`. So `3^3 | 9!`
- choosing bigger of `8` and `9` - in this case it is `9`
- so `s(432) = 9`

This solution runs around a minute.

## Alternative solutions

Consider `n = p1^a1 * p2^a2 * ... * pk^ak`

Then `s(n) = max(s(p1^a1), s(p2^a2), ... s(pk^ak))`

Calculation of `s(p^a)` can be done using the method from original solution.

## External links

https://www.geeksforgeeks.org/prime-factorization-using-sieve-olog-n-multiple-queries/

https://mathworld.wolfram.com/SmarandacheFunction.html