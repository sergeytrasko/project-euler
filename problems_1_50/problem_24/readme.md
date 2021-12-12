# Project Euler Problem 24 - Lexicographic permutations

[https://projecteuler.net/problem=24](https://projecteuler.net/problem=24)

## Solution idea

I know that there exists an algorithm to find N-th lexicographic permutation, but I don't know how it works.

So I decided to go with naive implementation. I know how to code recursive permutations generator. It will give us 10! (around 3.5 million) permutations. Then I can sort them using golang's `sort` package and take 999'999'999-th element (remember that arrays start with 0).

## Alternative solutions

The proper solution is based on so called factoradic numbers (numbers in factorial number system). That allows to compute N-th permutation in linear time.

## External links

https://en.wikipedia.org/wiki/Factorial_number_system
