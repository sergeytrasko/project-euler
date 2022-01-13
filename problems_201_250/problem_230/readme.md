# Project Euler Problem 230 - Fibonacci Words

[https://projecteuler.net/problem=230](https://projecteuler.net/problem=230)

## Solution idea

Concatenation of `A` and `B` sequences work like Fibonacci numbers.

The solution for `D(A, B, n)` can be found in 2 steps

Step 1 - understand how many times `A` and `B` should be concatenated. Basically it's finding `i`-th Fibonacci number with `f(0) = len(A)` and `f(1) = len(B)`.

Once we found `i` we can backtrace and see where `n`-th digit comes from - from `f(i-1)` or `f(i-2)`. This way we can go back to the initial `A` and `B` strings.

We need to run this algorithm for each of 18 digits and concatenate them properly.

## Alternative solutions

## External links
