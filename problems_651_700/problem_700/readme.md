# Project Euler Problem 700 - Eulercoin

[https://projecteuler.net/problem=700](https://projecteuler.net/problem=700)

## Solution idea

I tried to brute-force and could generate first 16 numbers in reasonable time. The last yielded Eulercoin was `15806432` and `n` was `42298633`.
Actually next Eulercoin would apear after ages.

That's why I needed different approach.

So far I know first 16 Eulercoins. And what if I calculate it backwards?

Notice that `a = 1504170715041707` and `m = 4503599627370517` are co-prime. This means that there exists inverse modular multiplicative - `a*aInv = 1 (mod m)` that is `aInv = 3451657199285664`. It means that the last Eulercoin `1` would be reached at `n = 3451657199285664`.

Basically what we need to do is just go from `1` to `15806432` (last Eulercoin found in normal way) and find Eulercoins based on inverse `a`.

## Alternative solutions

## External links
