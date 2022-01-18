# Project Euler Problem 304 - Primonacci

[https://projecteuler.net/problem=304](https://projecteuler.net/problem=304)

## Solution idea

Fibonacci number for big values of `n` can be calculated in `O(logn)` using recursive formula with memorization.

And then we can iterate from `10^14` and find probably prime numbers using a probability primality test (built-in golang package) based on Miller-Rabin test (that works reliably on numbers below `2^64 > 10^14`)

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Fibonacci_number