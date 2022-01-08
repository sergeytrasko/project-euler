# Project Euler Problem 160 - Factorial trailing digits

[https://projecteuler.net/problem=160](https://projecteuler.net/problem=160)

## Solution idea

Naive implementation works well for `n < 10^9` (or maybe even `10^10`) - just multiply from `1` to `n`, take results `mod 10^10` and during the multiplication process truncate trailining zeros.

But this will not work for `n = 10^12`.

Trailing zeros in factorial come of `2` and `5` factors. Basically `n!` can be represented as `2^a * 2^b * x` (`gcd(2, x) = gcd(5, x) = 1`) and number of trailing zeros will be `min(a, b)`. So there should be probably some trick with factors of `2` and `5`.

Let's do some exploration - take smaller modulo (e.g. `100` instead of `100000`) and calculate some first values of `f(a * 100)` to see any pattern.

What I noticed is that `f(100!) = f(500!) = f(f2500!)`, similarly `f(200!) = f(1000!) = f(5000!)`.

Similar pattern applies for `f(100'000!) = f(500'000!)` (if we take last 5 digits). But after additional exploration it turns out that this relation only works if `n = 0 (mod 2500)`

It gives a hypothesis: if `k >= mod` then `f(k) = f(5*k)`.
If it is so then `f(10^12!) = f((10^12 / 5^k) !)` and we need to have `10^12/5^k > 100'000`.

We can guess or check that `k <= 10`. But also `10^12/5^k = 0 (mod 2500)` and in such case `k <= 8`

So `f(10^12!) = f((10^12 / 5^8) !) = f(2560000!)`

Calcuating `2560000!` is rather straight-forward and can be done using naive implementation.

## Alternative solutions

You can cheat and use Wolfram Alpha for the solution - https://www.wolframalpha.com/input/?i=%2810%5E12%29%21

## External links
