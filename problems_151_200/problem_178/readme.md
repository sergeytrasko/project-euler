# Project Euler Problem 178 - Step Numbers

[https://projecteuler.net/problem=178](https://projecteuler.net/problem=178)

## Solution idea

Very similar idea and approach as for [problem 162 - Hexadecimal numbers](../problem_162).

I used dynamic programming to solve it.

Consider `dp[n][m][d]` 3-dimensional array where:
- n = 40 - number of digits
- m = 1024 - bitwise mask showing if digit 0, 1, ... 9 are present in the number
- d = 10 - last digit of the number

Bitwise mask works in the following way: if there was a number that used digits 1, 3 and 8 the bitwise mask `m` will be `2^1 + 2^3 + 2^8 = 2 + 8 + 256 = 266`

On every step we calculate how many numbers of length `n` with last digit `d` are there:
- if `d > 0` then it will be same as number of numbers with length `n-1` and where last digit was `d-1`. Also keeping track of used digits as bitwise mask
- similarly for `d < 9` - last digit is `d+1`

Final answer will be sum across all `n` and `d` of `dp[n][1023][d]` (1023 means that all 10 digits are used)

## Alternative solutions

There is no other viable solution - no way to go though 10^40 numbers and check every of them.

With dynamic programming approach the number of operations we need to do is around `40 * 1024 * 10` (40 digits number, 1024 options in bitwise mask and 10 digits)

## External links

https://en.wikipedia.org/wiki/Dynamic_programming