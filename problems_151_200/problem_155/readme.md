# Project Euler Problem 155 - Counting Capacitor Circuits

[https://projecteuler.net/problem=155](https://projecteuler.net/problem=155)

## Solution idea

The problem can be solved using recursion.

Define `solve(n)` function that returns all possible capacitances for a `n` capacitators.

We can split `n` into `a` and `b` such that `a+b<=n` and `a<=b`.

Then we can recursively calculate `solve(a)` and `solve(b)` and try to connect each pair in sequence and in parallel.

On each step eliminate duplicates.

Important implementaiton aspects:
- I suspect that using `float64` for calculating capacitance values might not be accurate. That's why I implemented `frac` structure and working with it
- `sequential` and `parallel` functions are useful for working with `frac` structure
- don't forget to use `frac` in lowest terms (i.e. reduce it)
- capacitators value does not matter

With all this in mind the solution runs around a minute.

## Alternative solutions

## External links

https://oeis.org/A153588