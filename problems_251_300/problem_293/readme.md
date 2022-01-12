# Project Euler Problem 293 - Pseudo-Fortunate Numbers

[https://projecteuler.net/problem=293](https://projecteuler.net/problem=239)

## Solution idea

We can generate all admissible numbers easily:
- just all powers of `2`
- recursively try to multiply sequential primes (don't forget about powers of primes). Note that number `n` should be a multiple of at least 2 sequential primes `p` and `q`. Therefore `p < q < sqrt(N)`.

Luckly there are not many such numbers - only 6656

We can check every of them and try to iterate over `n+m` and test each of them if it is a prime number or not (using probability test) and collect distinct values of `m`.

## Alternative solutions

## External links
