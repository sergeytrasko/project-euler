# Project Euler Problem 188 - The hyperexponentiation of a number

[https://projecteuler.net/problem=188](https://projecteuler.net/problem=188)

## Solution idea

Hyperexponentiation is also called Tetration. For notation we will use `T(a, b) = a^^b` 

Basically we need to to calculate `1777^^1855 mod 10^8 = T(1777, 1855) mod 10^8`.

Hint to solving the problem is the following statement:
- if `gcd(a, m) = 1` then `a^b mod m = a^(b mod phi(m)) mod m` where `phi(m)` is Euler's totient function.

In our case `a` is 1777 and `m` is 10^8 - and obviously `gcd(a, m) = 1`.

Basically there is a recursive relation: `T(a, b) mod m = a^T(a, b-1) = a^T(a, (b-1) mod phi(m)) mod m`

Some special cases:
- `m = 1` - in this case `T(a, b) = 0` as `x = 0 mod(1)`
- `b = 1` - in this case `T(a, 1) = a` as `T(a, 1) = a^T(a, 0) = a^1 = a`

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Tetration
https://en.wikipedia.org/wiki/Euler's_totient_function
https://stackoverflow.com/questions/30713648/how-to-compute-ab-mod-m
