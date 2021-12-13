# Project Euler Problem 111 - Primes with runs

[https://projecteuler.net/problem=111](https://projecteuler.net/problem=111)

## Solution idea

Looking at example provided for the problem I made the following assumptions:
- Seems that `M(n, d)` should be close to `n-1`. In other words resulting primes will contain a lot of same digits
- As consequence, `N(n, d)` should be relatively small - there should be not much primes that have repeating digits

First step is to create a function that can generate integers that have defined number of digits `d`. I created recursive function `generate`. What it does is just tries to append every possible digit (handling leading 0 of course) and keeps track of `sameDigitCount`.
As expected for high values of `M(n, d)` there are no many such numbers.

Once we have numbers with repeating digits generated we need to filter out only primes. As our limit is 10-digit numbers it is sufficient to check if every generated number is divisible by any prime below `10^5`.

Surprisingly my assumptions were correct, and `M(n, d)` was `9` or `8` leaving us with handful of primes to sum.

In reality the problem is easier than it seemed at first sight.

## Alternative solutions

## External links
