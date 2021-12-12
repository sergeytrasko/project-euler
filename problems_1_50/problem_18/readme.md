# Project Euler Problem 18 - Maximum path sum I

[https://projecteuler.net/problem=18](https://projecteuler.net/problem=18)

## Solution idea

Another classic dynamic programming problem.

To each element (x, y) we can come in 2 ways - from one of two elements above, i.e. (x, y-1) or (x-1, y-1).
Then max path to element (x, y) can be calculated as `element(x, y) + max{element(x, y-1), element(x-1, y-1)}`.

Note that half of the solution's code just reads the data from input file.

## Alternative solutions

Recursive solution also should work (as mentioned in the problem statement there are 16384 routes), however we should keep in mind also [problem 67](../../problems_51_100/problem_67) that has much higher number of routes.

## External links

https://en.wikipedia.org/wiki/Dynamic_programming
