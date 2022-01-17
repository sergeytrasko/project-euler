# Project Euler Problem 618 - Numbers with a given prime factor sum

[https://projecteuler.net/problem=618](https://projecteuler.net/problem=618)

## Solution idea

Entering first numbers into Open Encyclopedia of Integer sequences gives us nice recursive formula that can be converted to dynamic programming solution.

However there is another way how to derive this formula:
- consider that we have some sum `S(n)` that have numbers `a` and `b`
- what happens if we add `p` as a new divisor? `a*p` and `b*p` will be part of `S(n+p)`, or `S(n) = S(n-p)*p` for all `p`
- we need to be careful here and should avoid double-counting - therefore we need to make sure that we are using only values of `p` that were not used before

And as we need to take last 9 digits we need to do all our calculations `mod 1e9`

## Alternative solutions

Almost the same solution as [problem 31 - Coin sums](../../problems_1_50/problem_31)

## External links

https://oeis.org/A002098