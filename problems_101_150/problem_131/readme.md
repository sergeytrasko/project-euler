# Project Euler Problem 131 - Prime cube partnership

[https://projecteuler.net/problem=131](https://projecteuler.net/problem=131)

## Solution idea

First idea is to try to brute-force for some smaller values of prime `p`.
It turns out pretty quickly that `n` should be a perfect cube! I cannot prove it, but I strongly believe it is so.

There are around 80'000 prime numbers below 1 million so the idea might be to iterate over all such primes and try to find corresponding `n = k^3` (within reasonable `k` let say between 1 and 1000). But it works too slow.

Next thing I noticed is that values of `k` are increasing and they are almost sequential. Meaning that if some primes `p1` and `p2` are sequential solutions then corresponding values of `k1` and `k2` differ by very small amount (around 5 max).
So I'm making assumption if we check some prime `p` and we know that for previous prime we had a solution with some `k_0`, then we need to check values of `k` between `k_0` and `k_0 + 20` (20 is a magic constant).

Surprisingly this approach worked.

Also as the numbers of `n` and `p` go way beyond 64 bits (especially `n^3`) we need to use golang's `big` package.
In the links section there is a reference how to calculate cubic root in golang (there is no such feature available).

## Alternative solutions

## External links

https://stackoverflow.com/a/51390715 - cubic root in golang
https://pkg.go.dev/math/big