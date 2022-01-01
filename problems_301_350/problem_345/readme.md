# Project Euler Problem 345 - Matrix Sum

[https://projecteuler.net/problem=345](https://projecteuler.net/problem=345)

## Solution idea

The problem can be rephrased as following:
- There are N different jobs
- There are N different workers
- Each worker can do any job for a certain price
- Find the assignment where the total price is the biggest,

This is an assignment problem and it can be solved using Hungarian algorithm.

Trick is that assignment problem is solved for minimization of the cost, but here we need to have the maximal cost.

I used external module that contains implementation of the algorithm.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Hungarian_algorithm
https://github.com/oddg/hungarian-algorithm
