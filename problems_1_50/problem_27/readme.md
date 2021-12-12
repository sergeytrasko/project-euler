# Project Euler Problem 27 - Quadratic primes

[https://projecteuler.net/problem=27](https://projecteuler.net/problem=27)

## Solution idea

Another problem where (it seems) there is no better solution just a brute-force.
As a and b are in range between -1000 and 1000 we can just iterate all over them (that gives us 4 million combinations) and calculate the `n^2 +an + b` sequence hoping that it will yield a composite number soon.

## Alternative solutions

Function for checking if a number is prime is not optimal. Better approach would be to generate prime numbers below some limit (I assume 1 million would be enough) and check if `n` is divisible by the prime numbers.

## External links
