# Project Euler Problem 207 - Integer partition equations

[https://projecteuler.net/problem=207](https://projecteuler.net/problem=207)

## Solution idea

Before coding we need to understand the relation between `t` and `k`.
Let's analyze the equation `4^t = 2^t + k`.

Make substitution `u = 2^t` that gives us quadratic equation = `u^2 - u - k = 0`.
Roots are: `u = (1 +- sqrt(1+4k))/2`. We can drop negatve branch as `sqrt(1+4k) > 1` and `u` should be positive by definition.
Making back substitution: `2^t = (1 + sqrt(1+4k))/2` or `2^(t+1) = 1 + sqrt(1+4k)`

Next observation: we need to have `1+4k` to be a perfect square. Notice if `1+4k = (2n-1)^ = 4n^2-4n+1` then `k = n(n-1)`
So we get that `2^(t+1) = 1 + 2n-1 = 2n` or `2^t=n`. In other words `k` should be written as `2^a*(2^a - 1)` for some positive integer `a`.

Again we can replace `2^a` with `x` and iterate through all `x` values and check (note that `2^a >=1` therefore `x >= 2`):
- if `x` is a power of `2` (i.e. `a` is an integer for `x = 2^a`) - this gives a new perfect partition
- otherwise it's just a partition
- keep in mind that `m >= k = 2^a*(2^a - 1) = x*(x-1) = x^2 - x`, so minimal `m` is reached for `x^2 - x`

We should stop when `perfect/total < 1/12345`.

## Alternative solutions

## External links
