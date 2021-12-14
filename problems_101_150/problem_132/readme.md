# Project Euler Problem 132 - Large repunit factors

[https://projecteuler.net/problem=132](https://projecteuler.net/problem=132)

## Solution idea

By definition `R(n)` is `11111...1111`. Or alternatively we can write `R(n) = (10^n-1)/9`.

If prime `p` divides `R(n)` then `(10^n-1)/9 = 0 (mod p)` or (as `p` is not equal to 9) we can write `10^n-1 = 0 mod (9*p)` or `10^n = 1 mod (9*p)`.

Golang's `big` package comes with `Exp` function that allows you to calculate `a^b mod m` - that's what we really need.

First of all we need a primes sieve - you can experiment a bit and make a guess that primes below a million should be sufficient.
And then for every prime `p` check if `10^(10^9) = 1 mod (9*p)` - this can be done with a few lines of code.

## Alternative solutions

According to wikipedia, if `gcd(m, n) = d`, then `R(m)` and `R(n)` will have also common factors from `R(d)`.
E.g. `R(10)` will have factors of `R(5)` and factors of `R(2)` (plus some others)
This gives us motivation to break down the problem into smaller ones.
For example factors of `R(10^9)` include factors of `R(2^9)` and factors of `R(5^9)` and then we can split `R(2^9)` and `R(5^9)` into smaller cases.
I tried this approach and unfortunatelly it did not work really well as:
- it's too slow
- After having divisors of `R(d)` you still need to factor `R(n)/R(d)` - there are chances that you will face a big prime

But never the less I think this approach can work for smaller `n`

## External links

https://en.wikipedia.org/wiki/Repunit