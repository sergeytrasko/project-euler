# Project Euler Problem 234 - Semidivisible numbers

[https://projecteuler.net/problem=234](https://projecteuler.net/problem=234)

## Solution idea

First I wrote a brute-force solution to get some intermediate results to check more optimal algorithm.

The fast solution is based on the following:
- generate all primes below `sqrt(999966663333) + <buffer>`
- note that if `lps(n)` and `ups(n)` are sequential primes
- this means that we can iterate primes instead of `n` and count how many numbers `n` between `p^2` and `q^2` are divisible only by `p` or only by `q`

Let's take example: `p = 3` and `q = 5`:
- This means that `9 < n < 25`
- The following numbers are divisible by `p = 3`: `12 = 3*3 + 3`, `15 = 3*3 + 2*3`, `18 = 3*3 + 3*3`, `21 = 3*3 + 4*3` and `24 = 3*3 + 5*3`
- The following numbers are divisible by `q = 5`: `20 = 5*5 - 5`, `15 = 5*5 - 2*5` and `10 = 5*5 - 3*5`
- There is a clear pattern - iterate `j` from `p^2 + p` and increase `j` by `p` unless `p^2` is reached
- Same in opposite direction for `q`
- Notice that we calculated `15` twice and it is divisible by both `p` and `q` - meaning it should be excluded from the sum twice

Surprisingly this approach works in milliseconds.

## Alternative solutions

## External links
