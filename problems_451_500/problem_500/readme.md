# Project Euler Problem 500 - Problem 500!!!

[https://projecteuler.net/problem=500](https://projecteuler.net/problem=500)

## Solution idea

Basic idea is to use divisor function.
If `x = p1^r1 * p2^r2 * ... * pk^rk` then number of divisors is `(1 + r1)*(1 + r2)*...*(1 + rk)`.
Obviously we can take first 500500 primes and have `x = p1 * p2 * ... * p500549 * p500500`, but the answer will be too big.

Explore some smaller values:
- for `16 = 2^4` divisors the answer is `2^3 * 3 * 5`
- for `32 = 2^5` divisors the answer is `2^3 * 3 * 5 * 7`
- for `64 = 2^6` divisors the answer is `2^3 * 3^3 * 5 * 7`

The pattern seems clear - we should have smaller primes witb bigger powers. But how big the powers are?

Imagine that we already know that for `2^k` divisors we have some `x = p1^r1 * ... * pm^rm`.
Let's take the next prime `q = p(m+1)` and decide what we should do. The options are:
- multiply `x` by `q` - this way we double the number of divisors
- or we can try to increase some of `ri` to `2*ri + 1`. Why so? we need to have `(1 + ri)` as a power of 2. Obviously `ri = 2^y -1` (e.g. 1, 3, 7, 15, ...). Then `2 * (2^y - 1) + 1` =  `2^(y+1) - 1` - also in the required format

To decide which option to choose we just evaluate what will give smaller result - either multiply `x` by new prime `q` or find some existing `pi` and check if "doubling" `ri` power will give us smaller result than `q`

To avoid integer overflows I'm using golang's `big` package.

The program runs around one minute.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Divisor_function