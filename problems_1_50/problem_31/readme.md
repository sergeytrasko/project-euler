# Project Euler Problem 31 - Coin sums

[https://projecteuler.net/problem=31](https://projecteuler.net/problem=31)

## Solution idea

I used dynamic programming approach to solve this problem.
Define array `values` containing values of all possible coins - 1, 2, 5, 10, 20, 50, 100, 200.
And then `p[n][k]` will represent a number of ways how `n` can be represented using at most `k` first different coins.

The result will be `p[200][8]`

It's nice to sit with pen and paper and model the algorithm for smaller number of `n` and `k`, e.g. for `n` equal to 8 and using only 1, 2 and 5 as coins.

## Alternative solutions

## External links
