# Project Euler Problem 401 - Sum of squares of divisors

[https://projecteuler.net/problem=401](https://projecteuler.net/problem=401)

## Solution idea

I found a nice formula on Online Encyclopedia of Integer Sequences that is based on Pyramidal numbers (I don't fully understand why this formula works):

`f(n) = n*(n+1)*(2*n+1)/6 = 1^2 + 2^2 + 3^2 + ... + n^2`

Also define `s = sqrt(n)`

Then `SIGMA2(n) = Sum{k=1..s} (f(n/k) + k^2*(n/k)) - s*f(s)`.

The most complex part is to put `% mod` in all places to avoid integer overflows

## Alternative solutions

## External links

https://oeis.org/A064602