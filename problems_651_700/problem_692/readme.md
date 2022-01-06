# Project Euler Problem 692 - Siegbert and Jo

[https://projecteuler.net/problem=692](https://projecteuler.net/problem=692)

## Solution idea

Looking at `G(23416728348467685)` it's clear that brute-force approach will not work here and there should be some trick. It should be either a closed form expression or a recursive formula that converges in logarthmic time.

As first step I created a recursive function `alwaysWins(remaining, canTake)` that calculates if a player always wins (irrespectively how opponent plays) if `remaining` pearls are left and on the next move player can take up to `canTake` pearls.

After observing first 50 results I noticed that the values for `h(n)` consists only of Fibonacci numbers.

Quick googling showed that there is a Fibonnaci Nim game with exactly same rules!

Winning strategy for heap of `n` pearls is simple - take the smallest term in Zeckendorf representation of `n`.

So the answer is sum of smallest terms of Zeckendorf representation from `1` to `23416728348467685` (that is a Fibonacci number as well).

Now we need to understand how to calculate it efficiently.

If we print first 35 values fo `h(n)` then we can observe the pattern (here `fib(k)` is k-th Fibonacci number):
- take `h(n)` values between `n=1` and `n=7` - they are the same as between `n=14` and `n=20`
- in more general case - `h(n)` on interval `(fib(k-1), fib(k) )` is the same as on interval `( 0, fib(k-2) )`

This gives us the following recursion: `G(fib(k)) = G(fib(k-1)) + G(fib(k-2)) - fib(k-2) + fib(k)`

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Fibonacci_nim
https://en.wikipedia.org/wiki/Zeckendorf%27s_theorem
https://oeis.org/search?q=23416728348467685
https://oeis.org/A139764