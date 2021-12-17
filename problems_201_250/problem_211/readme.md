# Project Euler Problem 211 - Divisor Square Sum

[https://projecteuler.net/problem=211](https://projecteuler.net/problem=211)

## Solution idea

Here we need to use some properties of divisor function.
The sum of divisors squares of `n` is also refered to `sigma2(n)` and it has the following properties that we will use:
- if `gcd(a, b) = 1` then `sigma2(a * b) = sigma2(a) * sigma2(b)` (multiplicative rule)
- there is a nice formula for `sigma2(n)` based on prime factorization: `sigma2(n) = mul ( (p^(a+1)*2 - 1)/(p^2-1) )`

That gives us a hint - we can generate all primes numbers `p` below 64'000'000 (there will be around 3 and a half million of them) and calculate `sigma2` for all `p`, for all `p^n` and for all `k*p^n` (keeping in mind that `gcd(k, p) = 1`) using multiplicative rule.

The problem with the solution is that it runs a bit slow, but still within 1 minute on my laptop

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Divisor_function#Properties