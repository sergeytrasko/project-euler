# Project Euler Problem 30 - Digit fifth powers

[https://projecteuler.net/problem=30](https://projecteuler.net/problem=30)

## Solution idea

For this problem we can iterate numbers from 1 to some limit. Then each number should be split into digits and calculate sum of the digits.

The biggest question is the limit. I chose 10'000'000 as a reasonable assumption.

In reality: imagine that we have m-digit number X. So `X <= 10^m`. Also maximal digits sum is achieved when every digit is 9. So fith power digit sum is less than `m*9^5`.
Now we need to solve inequality: `10^m > m*9^5`, e.g. using [Wolfram Alpha](https://www.wolframalpha.com/input/?i=10%5Em+%3E+m*9%5E5) and we get that `m` is between 5 and 6. So basically the threshold can be lowered to 1'000'000

## Alternative solutions

## External links
