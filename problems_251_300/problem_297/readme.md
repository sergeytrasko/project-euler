# Project Euler Problem 297 - Zeckendorf Representation

[https://projecteuler.net/problem=297](https://projecteuler.net/problem=297)

## Solution idea

Problem states that we should sum the values of `z(n)` function with n up to `10^17` - that's obviously not possible and there should be some recursion or closed form for this expression.

I started with exploration and wrote implementation for `z(n)` function and `s(n)` function that just sums `z(i)` for `i=1..n` to get impression of the values.

Next I looked at this image from Wikipedia: https://en.wikipedia.org/wiki/Zeckendorf%27s_theorem#/media/File:Zeckendorf_representations_89px.svg

Based on it I made the conclusion: if `n = fib(k)` (i.e. input is a Fibonacci number) then we can compose `s(n) = s(fib(k))` as:
- if we remove bigger rectangle (that is `fib(k)`) then remaining part looks like `S(fib(k-1))`
- plus we need to add `S(fib(k-2))`

So the recurrent formula looks like `S(fib(k)) = S(fib(k-1)) + (fib(k) - fib(k-1)) + S(fib(k-2))`.

This works only for `n` that are Fibonacci numbers. Obviously `10^17` is not a Fibonacci number.

But that should not be a problem, as we can say (consider that `k` is the biggest integer such that `fib(k) <= n`) that `S(n) = S(fib(k)) + (n - fib(k)) + S(n-fib(k))`.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Zeckendorf%27s_theorem