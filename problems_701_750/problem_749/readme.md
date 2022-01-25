# Project Euler Problem 749 - Near Power Sums

[https://projecteuler.net/problem=749](https://projecteuler.net/problem=749)

## Solution idea

First thing I did - explore the small values.

It turns out that the 2 values for `S(2)` are `35` and `75` and they both have `k = 2`. Similarly for `S(6)` ther are 6 values: two of them ar `35` and `75` and other 4 are 6-digit numbers with `k = 6`.

Exploring slightly bigger cases I noticed that seems that `k = n` for `n`-digit numbers. At least it was for `S(8)`.

Let's define bounds for `k`.

If number `X` has `n` digits, then `10^(n-1) < X < 10^n`. Similarly min digit power sum is `1` (if `X = 10000...000`). Max digit power sum is `n*9^k` (if `X = 99..9999`).

Define upper bound for `k` like this:

```
n*9^k < 10^n
lg(n) + k*lg(9) < n
k < (n - log(n))/lg(9)
```

Looking at the data we can conclude that `k <= n`.

Next important observations: imagine that 2 numbers `x` and `y` consist of same digits `a`, `b` and `c`, e.g. `x = abc` and `y = bca`. Then their power digit sum will be the same.

This leads to solution idea:
- generate all possible digits combinations for digit sum from `1` to `9*16`
- calculate all power digit sums for `k` from `2` to `16`
- check if power digit sum consits of same digits as were generated.

Implementation for digits split is taken from [problem 171 - Finding numbers for which the sum of the squares of the digits is a square](../../problems_151_200/problem_171)

## Alternative solutions

## External links

https://oeis.org/A300160