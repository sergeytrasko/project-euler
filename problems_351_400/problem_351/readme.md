# Project Euler Problem 351 - Hexagonal orchards

[https://projecteuler.net/problem=351](https://projecteuler.net/problem=351)

## Solution idea

Few observations:
- orchard can be splited into 6 parts (they are symetric)
- number of points on `k`-th layer of 1/6-th of orchard is `k`
- we can make a square grid out of triangular grid - it is easier to work with and notice the relations
- each point can be represented with coordinates `(p, q)`

Important claim: point `(p, q)` (located on `p`-th layer) is visible if and only if `p/q` is a proper fraction (or `gcd(p, q) = 1`).

That means that the number of visible points on `p`-th layer is `phi(p)` (Euler's totient function).

But then the number of hidden points is `p - phi(p)`.

So the answer is `H(n) = 6 * sum{k = 1..n} (k - phi(k)) = 6*n*(n+1)/2 + sum{k = 1..n} (phi(k)) = 3*n*(n+1) + sum{k = 1..n} (phi(k))`

There is a nice way how to calculate sum of `phi` functions (see Totient summatory function) using Mebius inversion.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Euler%27s_totient_function

https://en.wikipedia.org/wiki/Totient_summatory_function

http://oeis.org/A064018