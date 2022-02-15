# Project Euler Problem 531 - Chinese leftovers

[https://projecteuler.net/problem=531](https://projecteuler.net/problem=531)

## Solution idea

Problem title gives a hint that it can be solved using Chinese Remainder Theorem, however it works only for co-prime moduli.

There is a generalization for non-coprime moduli `m` and `n`.

Let `g = gcd(m, n)`.

If `a = b (mod g)` then there is a solution.

We can use extended Euclide's algorithm to find such `u` and `v` so that `g = u*m + v*n`.

Then `x = (a*v*n + b*u*m)/g`. We just need to ensure that `x` is between `0` and `m*n/g`.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Euler%27s_totient_function

https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Generalization_to_non-coprime_moduli