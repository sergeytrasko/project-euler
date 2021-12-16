# Project Euler Problem 78 - Coin partitions

[https://projecteuler.net/problem=78](https://projecteuler.net/problem=78)

## Solution idea

My first idea was to use dynamic programming approach similar to [problem 31](../problem_31), but after calculating to `N=100` and getting quite big numbers and not noticing any ending zeros, I realized that the answer is much bigger and result will contain a lot of digits.

Quite quickly I found information about Pentagonal theorem and it's relation to number of integer partitions.
It turned out that there is a recurrent formula for number of partitions:
`p(n) = p(n-1) + p(n-2) - p(n-5) - p(n-7) + p(n-12) + p(n-15) - p(n-22) - p(n-26) + ... = sum( (-1)^k * p(n - g(k)) )`
Where each of the terms are generalized pentagonal numbers: `g(k) = k(3k-1)/2` for `k = 1, -1, 2, -2, 3, -3, ...`

It can be expected that the answer `p(n)` will be quite large and we should use golang's `big` package, however we can do a trick her and just calcuate `p(n) mod 1'000'000`. This will allow us to use 32 or 64 bit integers in golang's native way.

## Alternative solutions

Dynamic programming is a possible option if we do it `mod 1'000'000`. However knowing the answer to the problem there will be a lot of memory required (around 10Gb) to solve the problem.

## External links

https://en.wikipedia.org/wiki/Pentagonal_number_theorem#Relation_with_partitions
