# Project Euler Problem 365 - A huge binomial coefficient

[https://projecteuler.net/problem=365](https://projecteuler.net/problem=365)

## Solution idea

First idea was to create a brute-force naive solution - I knew that it will not solve the problem, but it would help to troubleshoot my optimal algorithm.

Quite quickly I came across Lucas's theorem that shows how to find a reminder of a binomial coefficient.

However it works only for prime modulo. 

But we know that the modulo we use `p*q*r` is a combination of 3 primes - this means we can use Chineese Reminder Theorem to solve 3 conguences.

In the end I found a blog post that outlines the exact same solution.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Binomial_coefficient

https://en.wikipedia.org/wiki/Lucas's_theorem

https://en.wikipedia.org/wiki/Chinese_remainder_theorem

https://fishi.devtail.io/weblog/2015/06/25/computing-large-binomial-coefficients-modulo-prime-non-prime/

