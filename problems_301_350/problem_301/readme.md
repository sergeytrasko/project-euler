# Project Euler Problem 301 - Nim

[https://projecteuler.net/problem=301](https://projecteuler.net/problem=301)

## Solution idea

Nim game is well analyzed from math perspecitive.

And it is proved that the position in winning if and only if `xor` operation between heap sizes (also referred as nim-sum) is 0 .

In other words `X(n, 2n, 3n) != 0 <=> n xor 2n xor 3n = 0`.

## Alternative solutions

Interesting fact is that the answer is a Fibonacci number.

## External links

https://en.wikipedia.org/wiki/Nim#Mathematical_theory