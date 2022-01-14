# Project Euler Problem 165 - Intersections

[https://projecteuler.net/problem=165](https://projecteuler.net/problem=165)

## Solution idea

The problem itself is not so complex:
- There are 5000 line segments
- We need to check all intersections - in total it will be `5000*4999/2` points
- There is a known algorithm to check if 2 segments intersect

The most compelex part is to deal with floating number rounding to detect distinct points. By trial and error approach I found that 12 fractional digits give correct result.

## Alternative solutions

## External links

https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect

https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line_segment