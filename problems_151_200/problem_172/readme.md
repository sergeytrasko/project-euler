# Project Euler Problem 172 - Investigating numbers with few repeated digits

[https://projecteuler.net/problem=172](https://projecteuler.net/problem=172)

## Solution idea

The problem can be solved using dynamic programming.
Define memorization array `dp[19][5^10]`:
- First dimension - length of the number
- Second dimension - how many times each of digit is used (can be 0, 1, 2, 3 or 4. 5 or more is the same as 4 - it does not affect the result)
- Value - how many such numbers with this digit allocation

On each step take each digit and sum number of numbers from previous steps that had less than 3 such digits (keeping track of used digits)

The solution runs around a minute, but I think the most part is spent on array manipulations (mainly creating new arrays)

## Alternative solutions

## External links
