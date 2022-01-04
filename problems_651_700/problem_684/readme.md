# Project Euler Problem 684 - Inverse Digit Sum

[https://projecteuler.net/problem=684](https://projecteuler.net/problem=684)

## Solution idea

First observation to make that `s(n)` will have a form of `A9...9` where `A = n mod 9` and number of `9` will be `n div 9`.

For example `s(29) = 2999` as `29 mod 9 = 2` and `29 div 9 = 3` (so there will be 3 nines and first digit is 2).

Closed form for `s(n)` is `(n mod 9) * 10^(n div 9) + 10^(n div 9) - 1 = 10^(n div 9) * ((n mod 9) + 1)-1`

Next step is to try to sum values of `s(n)` in some chunks. If you do it by 9 after 2-nd group (first are `45` and `531`) there will be a clear pattern:
```
sum(9k) + sum(9k+1) + ... sum(9k+9) = 539....91
```
And number of nines will be `k-2`.

We can write that `Chunk(k) = sum(9k) + sum(9k+1) + ... sum(9k+9) = (54 * 10^k - 1)*10 + 1 = 540*10^(k-1)-9`

Sum of first `m` chunks (that is `S(9m+9)`) is:
```
Chunk(0) + Chunk(1) + ... + Chunk(m) = 
  (540*10^(0-1)-9) + (540*10^(1-1)-9) + ... + (540*10^(m-1)-9) =
  540*(10^-1 + 10^0 + 10^1 + ... + 10^(m-1)) - (m+1)*9 = 
  54 + 540*(10^0 + 10^ + ... 10^(m-1)) - (m+1)*9 = 
  54 + 540 * (10^m - 1)/9 - (m+1)*9
```

Verify few results:
- `m = 0`: `S(9) = 54 + 540*(10^0-1)/9 - 1*9 = 54 + 0 - 9 = 45`
- `m = 1`: `S(18) = 54 + 540*(10^1-1)/9 - 2*9 = 54 + 540 - 18 = 576`. On the other hand: `S(18) = S(20) - s(20) - s(19) = 1074 - 199 - 299 = 576`

The rest is relatively simple - just sum the big chunk and add required `s(k)` to reach `n` (it will be at most 9 such calls to `s(k)`).

As the numbers are quite big we need to be careful with integer overflow and modular operations (especially exponents).

Interestingly that Fibonacci numbers do not play any role here.

## Alternative solutions

## External links

https://oeis.org/A051885
