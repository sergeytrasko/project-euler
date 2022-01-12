# Project Euler Problem 169 - Exploring the number of different ways a number can be expressed as a sum of powers of 2

[https://projecteuler.net/problem=169](https://projecteuler.net/problem=169)

## Solution idea

I started with brute-force approach like:
- coverted `10^25` to binary
- created recursive function that tried to replace `1` bit with combination of lower powers

This worked well up to around `10^18` but it gave me some first values.

I wasn't able to deduce any relation (nor find a pattern any helpful pattern except that `f(2^n) = n+1` and `f(2^n-1) = 1`) between the numbers and decided to generate first 100 values and search them in [On-line Encyclopedia of Integer Sequences](http://oeis.org/). To my surprise it led me to Stern's diatomic series that had a nice recurrent formula.

The rest was relatively easy to implement.

2 implementation details:
- need to use long arithmetics as `2^64 < 10^25`
- memorization is required to cache results of recursive calls

## Alternative solutions

## External links

http://oeis.org/A002487

https://mathworld.wolfram.com/SternsDiatomicSeries.html

https://en.wikipedia.org/wiki/Calkin%E2%80%93Wilf_tree#Stern's_diatomic_sequence