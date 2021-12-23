# Project Euler Problem 225 - Tribonacci non-divisors

[https://projecteuler.net/problem=225](https://projecteuler.net/problem=225)

## Solution idea

For original Fibonacci sequence there is a known fact that members of the sequence are periodic some modulo. This is also called Pisano period.

I thought that the same approach might work for Tribonacci numbers - and indeed it was the right guess.
For each odd number `x` (starting from 3) I calculate Tribonacci numbers `mod x`. The outcome might be:
- we get some `T(n) = 0 (mod x)` - this means that `x` divides some numbers in Tribonacci sequence
- we get `T(n-2) = T(n-1) = T(n) = 1` - this means we got a cycle `mod x` and no number in this cycle is divisble by `x` - this is a number we want

Just repeat this exercise for 124 successful attempts.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Pisano_period