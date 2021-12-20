# Project Euler Problem 757 - Stealthy Numbers

[https://projecteuler.net/problem=757](https://projecteuler.net/problem=757)

## Solution idea

Imagine that `N  = ab = cd` and `a + b = c + d + 1`.
Set `x = c - a` and `y = d - a`.
Notice that `xy = (c - a)(d - a) = cd - a(c+d) + a^2 = N - a(a+b-1) + a^2 = a`
Similarly `(x + 1)(y + 1) = xy + x + y + 1 = a + c - a + d - a + 1 = (c + d + 1) - a = a + b - a = b`

This means that `N = xy(x+1)(y+1)` for some positive integers `x` and `y`.

Also it turns out that these numbers are called `bipronic` numbers.

Considering that `x > y` means that we should check only `x < sqrt4(10^14)`

Important consideration - some stealthy numbers can be composed in multiple ways, but we need to count them only once. For this we need to use some kind of `set` or `map` structure. Unfortunatelly golang is quite slow with huge maps (haven't made benchmarks for other languages for similar data volumes) - 99% of execution time is spend on tracking which `n` were already marked as solution.
For me the solution runs around 1 minute.

## Alternative solutions

For smaller numbers (by my observations - for `n < 10^10`) the following approach can be used:
- iterate over `a`
- iterate over `b`
- notice that `n` should be a multiple of 4 - cannot prove it, but it seems so from looking at the first solutions
- then make the following logical chain: `c*d = N => d = N/c = > a + b = c + N/c + 1`
- this can be used to convert `(a+b)*c = c^2 + N + c` to quadratic equation `c^2 + c(1-a-b) + N = 0`
- we can solve for `c` and check if it is a positive integer

## External links

https://oeis.org/A072389