# Project Euler Problem 171 - Finding numbers for which the sum of the squares of the digits is a square

[https://projecteuler.net/problem=171](https://projecteuler.net/problem=171)

## Solution idea

Maximum sum of digit squares is `19 * 9^2 = 1539`. That means that there are only `39` possible squares - `1, 2^2=4, 3^2=9, 4^=16, ..., 39^2=1521`.

As first step we need to generate all possible ways how to arrange digits from `1` to `9` so that their sum of squares is also a perfect square. This can be done recursively.

Next observation: consider a number consisting of digits `a`, `b` and `c`. We can write it as `abc`. This number will have the same digit square sum as `bca` and `acb`. Basically permutation of digits does not change digit square sum.

Also adding `0` as digits does not change digit square sum. E.g. `abc` will have the same digit square sum as `a0000b00c0000000000`.

After we got all these ways to get a perfect squares using digits from `1` to `9` we need to do the following:
- add as many `0` as possible
- get all permutations of the number
- sum it all together
- get result modulo `10^9`

Actually there is a short formula how to get sum of all permutations of a number - see link attached.

## Alternative solutions

## External links

https://www.quora.com/What-is-a-formula-to-calculate-the-sum-of-all-the-permutations-of-a-given-number-with-repitations