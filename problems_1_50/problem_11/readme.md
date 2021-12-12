# Project Euler Problem 11 - Largest product in a grid

[https://projecteuler.net/problem=11](https://projecteuler.net/problem=11)

## Solution idea

Iterate through whole grid and for each cell check 4 directions:
- right
- down
- right-down
- left-down

Note that we are checking only 4 out of 8 directions - other 4 will be checked when processing other cells. This also explains why we are processing 16 by 16 grid (despite original size is 20 by 20).

## Alternative solutions

## External links
