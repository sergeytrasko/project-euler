# Project Euler Problem 243 - Resilience

[https://projecteuler.net/problem=243](https://projecteuler.net/problem=243)

## Solution idea

Resilience will be equal to `phi(d)/(d-1)` where `phi(d)` is Eulers totient function.
Furthermore, `phi(d) = d * (1-1/p1)*(1-1/p2)*...(1-1/pn)`, where `pi` - different prime factors of d.
To mimimize `phi(d)` (and in the same time - minimize resilience) we need to have more smaller prime factors, ideally `d` should be `2*3*5*7*...`.

The strategy is:
- generate first primes
- start with `d = 1`
- on each step multiply `d` by next prime until we break required ratio `15499/94744`
- after trial and error we can see that the result is too low, and probably we can do something better
- I tried to replace last prime factor with some smaller number (not necessary prime) - i.e. 2, 3, 4, ... - until we are again below our ratio

In the end `d = 2*3*5*7*11*13*17*19*23*4`

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Euler%27s_totient_function