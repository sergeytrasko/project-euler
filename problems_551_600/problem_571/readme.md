# Project Euler Problem 571 - Super Pandigital Numbers

[https://projecteuler.net/problem=571](https://projecteuler.net/problem=571)

## Solution idea

We can generate all permutations of `[0, 1, 2, ... , N-1]` in lexicographic order (see link below for the algorithm), convert every permutation to a number and check if it's a pandigital number in all bases from `2` to `N-1`.

This approach works in a couple of seconds for `N=10` and gives correct answer from the example.

However for `N=12` it needs to process over `N! = 12!` permutations (that's around half a billion). We can do micro-optimizations here (like start with `[1, 0, 2, ..., 11]` as first permutation and check bases from `N-1` to `2` as there are more chances for pandigital number in smaller bases).

In the end it gives a solution in a couple of minutes.

Eventually it seems to be the most pragmatic approach to solve this problem (and other solvers did similar solution).

## Alternative solutions

## External links

https://www.geeksforgeeks.org/lexicographic-permutations-of-string/