# Project Euler Problem 129 - Repunit divisibility

[https://projecteuler.net/problem=129](https://projecteuler.net/problem=129)

## Solution idea

The problem description is quite tricky, but it can be rephrased as: need to find smallest `n` over 1'000'000 such that `n` is a divisor of `R(k)`.
Another important observation (or a guess) is that `A(n) < n`. So basically we need to check every `n` after 1'000'000 and find what is the minimal `k` when `n` is a divisor of `R(k)`. Basically here we can work `mod n` to avoid long arithmetics.

Condition `GCD(n, 10) = 1` can be simplified to `n % 5 == 1 && n % 2 == 1` - this also makes the implementation easier.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Repunit