# Project Euler Problem 485 - Maximum number of divisors

[https://projecteuler.net/problem=485](https://projecteuler.net/problem=485)

## Solution idea

The problem boils down to 2 sub-problems:
- find number of divisors of all numbers from `1` to `100'000'000`
- find maximum values of a sliding window of an array

Sum of divisors can be found based on prime factorizations.

For sliding window subproblem there is an alogrihtm that runs in linear time and uses deque (double ended queue).

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Divisor_function

https://www.geeksforgeeks.org/sliding-window-maximum-maximum-of-all-subarrays-of-size-k/