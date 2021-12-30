# Project Euler Problem 168 - Number Rotations

[https://projecteuler.net/problem=168](https://projecteuler.net/problem=168)

## Solution idea

Consider number `A = 10 * x + d` that has `n` digits.

Then it's right-rotation `B = d * 10^(n-1) + x`. Also we know that `B` is a multiple of `A` or `B = A*k` where `k` is from 1 to 9.

Rewriting it all we get: `d * 10^(n-1) + x = k * (10 * x + d) = 10 * k * x + k*d`.

After some simplification: `x * (10*k - 1) = d * 10^(n-1) - k*d` or `(d * 10^(n-1) - k*d) mod (10*k - 1) == 0`.

We know that `d`, `k` and `n` are from fixed range - we can just iterate though them.

## Alternative solutions

I started with naive implementation to get correct result for small cases to verify my algorithm.

Next I implemented my algorithm using `int` and `int64` types and compared it naive implementation. 

And only then I reimplemented in using golang's `big` package.

## External links
