# Project Euler Problem 15 - Lattice paths

[https://projecteuler.net/problem=15](https://projecteuler.net/problem=15)

## Solution idea

I used dynamic programming approach - for every cell we calculate the soltuion based on the information from previous steps.

In this particular case - for cell (x, y) the number of path is equal to sum of paths from left and from top.

## Alternative solutions

It turns out that it can be solved with combinatorics and the result is equal to 2N choose N (binomial coefficient - `C(2N, N)`).

The recursive solution will not work here as the number of paths is way too big and it goes beyond what a modern computer can calculate.

## External links

https://en.wikipedia.org/wiki/Dynamic_programming
https://en.wikipedia.org/wiki/Binomial_coefficient
