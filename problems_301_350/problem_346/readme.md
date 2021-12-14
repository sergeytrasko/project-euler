# Project Euler Problem 346 - Strong Repunits

[https://projecteuler.net/problem=346](https://projecteuler.net/problem=346)

## Solution idea

Repunit of length `n` and base `b` is equal to `(b^n-1)/(b-1)`
First of all we need to find all repunits in some base `b`.
To do it we iterate over all `b` from `2` to infinity and then through `n` from `3` to infinity.
We exit the loop when generate repunit is exeeding the limit of `10^12`.
Important aspect - why iterate `n` from `3` and not from `2` (as there might be repunits of length `2`)?
Imagine we found `r = repunit(n, b)` and we want to find other repunit with same value in different base.
Assuming that we want to find other repunit with bigger base. Obviously such repunit should have less digits, but at least 2.

How do we fing other repunit. We know that `r` is a repunit with `n` digits.
What we need to to is to check if there are **integer** solutions to `r = (b^a-1)/(b-1)` for all `a` between `2` and `n-1`
Notice that `r - (b^a-1)/(b-1)` is an increasing function (with fixed `a`) - it can be shown by expanding the expression - therefore should have exactly one root. The question is if it is integer or not.
This can be done with bisection method in `O(log(n))` complexity.

Things to consider:
- `1` is a strong repunit
- Don't count same strong repunit multiple times. E.g. `31` can be written as `11111_2`, `111_5`, `11_30`, but it should be counted only once in the target sum

## Alternative solutions

Another approach would be to iterate through all `n` and `b` (as in original solution) and fill in the map of repunits (key - repunit, value - number of times it was generated) and then filter out the ones that were present at least twice

## External links

https://en.wikipedia.org/wiki/Repunit
https://en.wikipedia.org/wiki/Bisection_method