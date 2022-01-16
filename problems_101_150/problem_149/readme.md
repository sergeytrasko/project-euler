# Project Euler Problem 149 - Searching for a maximum-sum subsequence

[https://projecteuler.net/problem=149](https://projecteuler.net/problem=149)

## Solution idea

Let's solve simplier version of the problem first - find maximum sub-array sum.

This can be done in linear time using Kadene's algorithm.

Once we have it implemented we can solve the original problem like:
- find max sum for every of 2000 rows
- find max sum for every of 2000 columns
- find max sum for every of 4000 diaginals
- find max sum for every of 4000 anti-diaginals

Take the largest value.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Maximum_subarray_problem