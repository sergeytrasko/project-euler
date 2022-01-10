# Project Euler Problem 249 - Prime Subset Sums

[https://projecteuler.net/problem=249](https://projecteuler.net/problem=249)

## Solution idea

Dynamic programming problem.

Define:
- `s` as array of primes below 5000, can be generated using primes sieve
- `ways[n]` as number of ways to compose `n` as sum of some primes
- `max` as sum of all elements of `s`

Inductive step is:
- for each element `p` from `s`
- for each `n` between `p` and `max` we have `ways[n] = ways[n] + ways[n-p]`

Once `ways` array is filled in - we just need to sum values of `ways[n]` where `n` is a prime below `max`. Again these primes `n` can be generated using sieve.

## Alternative solutions

## External links
