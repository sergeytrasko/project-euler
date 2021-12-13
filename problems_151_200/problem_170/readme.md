# Project Euler Problem 170 - Find the largest 0 to 9 pandigital that can be formed by concatenating products

[https://projecteuler.net/problem=170](https://projecteuler.net/problem=170)

## Solution idea

I used clever brute-force approach here. 
Some naming convention. Let `N` be an array of pandigital input numbers with `N0` is the first one.
This means that all products of `M1 = N0*N1`, `M2 = N0*N2`, ..., `Mk = N0*Nk` are pandigital.

The problem can be split in several sub-problems.
- Generate all pandigital numbers of length `l` using a set of digits (let's call it `d`)
  - As first substep - we can generate all combinations of size `l` of `len(d)`
  - For each combination compose all permutations
  - Convert permutation to a number, e.g. `[1, 2, 5] -> 125`. Note that number cannot start with 0
- Recursively generate `N` array starting with `N0`
- On each recursive step try to find `i`-th number `Ni` out of still not used digits and make sure that `N0*Ni` uses still unused digits
- Once we get a combination of `N` that is 0-9 pandigital and also `M` array also is 0-9 pandigital we need to find the best concatenation
  - We just need to concatenate `M` elements sorted by first digit - this will maximize the concatenation


Unfortunatelly I have a bug in my code that generates duplicates for pandigital numbers - that forces my algorithm to check same options multiple times.

Another optimization is that `N0` cannot have more than 5 digits - problem states that there should be at least 2 `N1`, `N2`, ... `Nk`, each of them at least 1 digit. So it case `k=2`, `N1` and `N2` will be already at least 5+5=10 digits long.

And also mu solution runs way too long - over 5 minutes, but gets the right answer.

## Alternative solutions

## External links
