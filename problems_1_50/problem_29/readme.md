# Project Euler Problem 29 - Distinct powers

[https://projecteuler.net/problem=29](https://projecteuler.net/problem=29)

## Solution idea

There will be in total 99*99 numbers of form `a^b`. Using golang's `big` package we can easily calculate `a` to power `b`.
What we are left to do - remove duplicates from calculated numbers. This can be done using `sort` package to sort the results and then iterate through the sorted array.

## Alternative solutions

## External links
