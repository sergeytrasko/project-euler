# Project Euler Problem 158 - Exploring strings for which only one character comes lexicographically after its neighbour to the left

[https://projecteuler.net/problem=158](https://projecteuler.net/problem=158)

## Solution idea

Let's simplify the problem like this:
- map letters `a`...`z` to `1`..`26`
- to calculate `p(n)` we take numbers from `1` to `n` and then multiply result of possible choices of `n` numbers (that is `C(26, n)`)

Thus the problem reduces to counting permutations where there is only one descent.

Quick googling routed me to Eulerian Numbers that has explicit formula.

The rest is easy.

## Alternative solutions

For small cases like `n=2`, `n=3` and `n=4` (that can be checked by hand) a recursion can be derived: `p(n+1) = 2^n-1+p(n)`

## External links

https://en.wikipedia.org/wiki/Eulerian_number

https://oeis.org/A000295