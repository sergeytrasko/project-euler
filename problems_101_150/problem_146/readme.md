# Project Euler Problem 146 - Investigating a Prime Pattern

[https://projecteuler.net/problem=146](https://projecteuler.net/problem=146)

## Solution idea

Some values of `n` can be excluded right away:
- odd `n` as `(2k+1)^2 + 1 = 4k^2 + 4k + 2` is an even number therefore not prime
- `n = 0 (mod 3)` as then `n^2+3` is divisible by 3
- Similarly `n` is not multiple of 7 and 13
- However `n` is multiple of 10 as:
  - we know that `n` cannot be odd therefore `n` is even
  - check `n mod 5` and see that for `n != 0 (mod 5)` any of `n^2 + k` is divisible by 5. For example, if `n = 3 (mod 5)` then `n^2+1 = 3^2 + 1 = 10 = 0 (mod 5)` meaning that `n^2+1` will not be a prime as it is divisible by 5

Basically this gives us a brute-force algorithm that works a bit slow, but produces the results.

I believe there are other ways how to limit number of cases to check.

## Alternative solutions

I tried to use an approach from [problem 216 - Investigating the primality of numbers of the form 2n^2-1](../problems_201_250/problem_216) with prime sieve for custom function. The approach is quite similar besides that we sieve 6 functions and 3 times more data.
However this works extremely slow for `n > 10'000'000`.

## External links
