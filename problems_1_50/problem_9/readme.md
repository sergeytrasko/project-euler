# Project Euler Problem 9 - Special Pythagorean triplet

[https://projecteuler.net/problem=9](https://projecteuler.net/problem=9)

## Solution idea

Brute-force solution. As `a+b+c=1000`, this means that `a`, `b` and `c` should be less than 1000 (you can do careful calculation and come up with better lower bound).
Just iterate a and b from 1 to 1000 and check if `a`, `b` and `1000-a-b` form a right triangle.

## Alternative solutions

Better solution would be to generate Pythagorean Triplets (e.g. using Euclid's formula) - the number of possibilities to check will be much less

## External links

https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple