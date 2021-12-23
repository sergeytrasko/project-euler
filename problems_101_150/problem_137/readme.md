# Project Euler Problem 137 - Fibonacci golden nuggets

[https://projecteuler.net/problem=137](https://projecteuler.net/problem=137)

## Solution idea

Fibonacci numbers can be expressed using a generating function:
```
A(x) = x / (1 - x - x^2)
```
And if setting `A(x) = k` and doing some algebra manipulations we come to a quadratic equation:
```
x = (-(k+1) + sqrt[ (k+1)^2 + 4k^2 ])/2k
```

Obviously for `x` to be a rational number `(k+1)^2 + 4k^2` should be a perfect square. 
That gives us an easy brute-force algorithm:
- iterate `k` from 1 to whatever we need
- calculate `s = (k+1)^2 + 4k^2`
- check if `s` is a perfect square

However this approach works only till `k = 11` until it hits integer overflow (even with 64 bit numbers). So different approach is needed.

Non-obvious mind-blowing observation (after searching online encyclopedia of integer sequences - link below) is that the answer is `F(2k)*F(2k+1)`!
Quote:
```
Nonnegative integers k such that G(x) = k for some rational number x where G(x) = x/(1-x-x^2) is the generating function of the Fibonacci numbers.
```

So 15-th Fibonacci nugget is `F(30)*F(31)`. So it boils down to calculating Fibonacci numbers till 31. Good thing is that they fit into 63 bits (and it's far beyond what is possible to do with brute-force approach)

## Alternative solutions

This problem is related to [problem 140 - ](../problem_140) that has a more logical solution using Diophantine equations.

## External links

https://en.wikipedia.org/wiki/Fibonacci_number#Generating_function
https://oeis.org/A081018
https://en.wikipedia.org/wiki/Generating_function
https://www.alpertron.com.ar/QUAD.HTM