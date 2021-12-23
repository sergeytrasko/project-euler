# Project Euler Problem 358 - Cyclic numbers

[https://projecteuler.net/problem=358](https://projecteuler.net/problem=358)

## Solution idea

First thing I did was a research about cyclic numbers.
It turned out that cyclic number of length `p` is related to digital period of `1/p`.
What are the bounds of `p`? If period of `1/p` starts with `00000000137` then it's fair to say that `1/p` is between `0.00000000137` and `0.00000000138` that gives us range of `p`: `724637681 < p < 729927007` that gives us around 5 million numbers that can produce such cyclic number.

Also we know that last digits are `56789`. Known property of cyclic numbers: if it is multiplied by generating number `p` then the result is all nines. Meaning that `p * 56789 = 999999 mod (100000)` - this should minimize number of candidates as well.

Actually very few numbers satisfy these conditions.

We can use adapted version of finding cycles of reciprocals from [problem 26 - Reciprocal cycles](../../problems_1_50/problem_26), but here we are interested only in digit sum that is nuch faster than calculating the actual cycle.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Cyclic_number