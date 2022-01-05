# Project Euler Problem 88 - Product-sum numbers

[https://projecteuler.net/problem=88](https://projecteuler.net/problem=88)

## Solution idea

Notation: let's define "`n` is minimal product-sum number for set size `k`" as function `mps(k) = n`

Few quick observations to define bounds of `mps(k)`:
- `mps(k) > k` - we can express `k = 1 + 1 + ... + 1`, but the product `1*1*1*...*1` obviously is less that `k`
- `mps(k) <= 2k` - we can express sum as `1 + 1 + ... + 2 + k = (k-2)*1 + 2 + k = 2*k` and in the same time `1*1*...*1*2*k = 2*k`. So `mps(k) = 2k` is a trivial solution.

Basically we need to check integers from `1` to `2*12'000 = 24'000` and check different factorizations and calulcate minimal `k` values.

Some optimizations:
- if the sum (or a product) of factors exceeds `2*n` - then it is a dead end and we will have a better solution
- on the last step (when `n` is already factored) - we just need to complete the sum with `1`-s

## Alternative solutions

## External links
