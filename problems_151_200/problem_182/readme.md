# Project Euler Problem 182 - RSA encryption

[https://projecteuler.net/problem=182](https://projecteuler.net/problem=182)

## Solution idea

Problem boils down to counting so called RSA fixed points.
It turns out that number of solutions for `m^e = m (mod n)` is equal to `(1 + gcd(e-1, p-1))*(1 + gcd(e-1, q-1)`

After it the solution becomes trivial.

Interesting fact - some papers mention that minimum number of RSA fixed points is 9 (that actually follows nicely form the equation above).

## Alternative solutions

## External links

https://math.stackexchange.com/questions/1298664/rsa-fixed-point
https://math.stackexchange.com/questions/58373/rsa-unconcealed-messages