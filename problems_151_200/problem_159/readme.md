# Project Euler Problem 159 - Digital root sums of factorisations

[https://projecteuler.net/problem=159](https://projecteuler.net/problem=159)

## Solution idea

Important fact is to notice the following relation: `mdrs(n) = max {drs(f) + mdrs(n/f)}` where `f` is a factor of `n` (except `1`).

We need to solve 3 sub-problems first:
- how to efficiently calculate digital root of a number. Actually it's pretty easy - it's just `drs(n) = (n-1) % 9 + 1` (see Wikipedia article about it).
- how to get all prime factors of a number
  - one approach would be to have a prime sieve and then then factor every number
  - better and faster approach is for each number calculate "smalles prime factor" - a smallest prime that divides number `n`. This can be calculated in linear time (not sure about if it is really `O(n)`, but for sure below `O(n^2)`). Then factorization of a number can be done in `O(logn)`.
- how to get every factor of `n` based on prime factorization - I used BFS approach here (similar to [problem 204 - Generalised Hamming Numbers](../../problems_201_250/problem_204))

The final algorithm is:
- calculate smallest prime factors for all numbers from `1` to `1'000'000`
- iterate from `1` to `1'000'000`
  - for each `n` calculate it's prime factorization
  - for prime factorization generate list of possible factors
  - for each factor calculate `drs(f) + mdrs(n/f)` (keep in mind that `n/f < n` and values of `mdrs(n/f)` is already calculated)
  - select max of the results and set it as value for `mdrs(n)`
- in the end sum all the results

## Alternative solutions

Even simplier solution is to use the fact that `mdrs` has some multiplicativity: for `n = a*b` (where `a <> 1` and `b <> 1`) we have that `mdrs(n) = max(max(mdrs(a) + mdrs(b)), drs(n))`

This approach even does not require prime factorization.

## External links

https://en.wikipedia.org/wiki/Digital_root
https://www.geeksforgeeks.org/prime-factorization-using-sieve-olog-n-multiple-queries/