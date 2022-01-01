# Project Euler Problem 133 - Repunit nonfactors

[https://projecteuler.net/problem=133](https://projecteuler.net/problem=133)

## Solution idea

Repunit `R(n)` can be expressed as `R(n) = (10^n-1)/9`. Similarly `R(10^n) = (10^(10^n)-1)/9`.

If `R(10^n)` has a prime factor `p` then `(10^(10^n)-1)/9 = 0 (mod p)` or `10^(10^n) = 1 (mod 9p)` (we can do it as `p` is not a multiple of 9).

Basically what we need to do is to check all values of `n` if `10^(10^n) = 1 (mod 9p)`. Notice that there will be at most `9p` distinct values of `10^(10^n) (mod 9p)` and there might be a cycle. There are around `9'500` primes below `100'000` so this approach seems feasible.

Obviously `10^(10^n)` can be a huge number and we can use a fact that `a^b mod m = a^(b mod phi(m)) mod m` where `phi` is Euler's Totient function.

It can be shown as `phi(9p) = phi(9)*phi(p) = 6 * (p-1)`.

My code runs in around a minute.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Repunit
https://en.wikipedia.org/wiki/Euler%27s_totient_function