# Project Euler Problem 193 - Squarefree Numbers

[https://projecteuler.net/problem=193](https://projecteuler.net/problem=139)

## Solution idea

Same solution as for [problem 745 - Sum of Squares II](../../problems_701_750/problem_745) (that I solved first...), but instead of sum just counting the numbers.

## Alternative solutions

Alternative solution implies usage of Moebius function that is defined as:
- `+1` if `n` is square-free integer with even number of prime factors
- `-1` if `n` is square-free integer with odd number of prime factors
- `0` if `n` has a factor of prime squared

Moebius function can be calculated in sieve-based approach for a range of `1..N`.

Then we just need to sum how many `i` are there such that:
- Moebius function of `i` is not `0`
- `i^2 <= N`

Pretty elegant solution that I (to be honest) do not understand 100% yet.

## External links

https://en.wikipedia.org/wiki/Inclusion%E2%80%93exclusion_principle

https://en.wikipedia.org/wiki/M%C3%B6bius_function