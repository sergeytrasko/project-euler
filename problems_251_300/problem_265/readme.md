# Project Euler Problem 265 - Binary Circles

[https://projecteuler.net/problem=265](https://projecteuler.net/problem=265)

## Solution idea

The problem is solved using recursion with bit arithmetics.

The recursive function keeps information about:
- current number
- used numbers (it will be `[2^n]bool` array)

On each step we do:
- shift current number by one bit right (divide it by 2)
- two branches exist:
  - append 0 - same as multiply by 2
  - append 1 - same as multiply by 2 and then add 1
- of course track records of which number we visited already

Finish when reached 0 again on the `2^n-1` step.

## Alternative solutions

## External links
