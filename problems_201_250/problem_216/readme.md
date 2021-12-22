# Project Euler Problem 216 - Investigating the primality of numbers of the form 2n^2-1

[https://projecteuler.net/problem=216](https://projecteuler.net/problem=216)

## Solution idea

Solution is based on idea of prime sieving.
For that we need few facts:
- if `t(n) = 0 (mod p)` then also `t(n+k*p) = 0 (mod p)`. This means that if we detected that `t(n)` is a composite with a factor `p` then `t(n+p)` is also composite. Proof: `t(n+k*p) = 2*(n+k*p)^2-1 = 2n^2-1+(4n*k*p+k^2*p^2) = t(n) + p*(4*n*k+k^2*p)`
- if `t(n) = 0 (mod p)` then also `t(-n) = 0 (mod p)`
- as consequence: if `t(n) = 0 (mod p)` then also `t(-n+k*p) = 0 (mod p)`

Algorithm:
- generate `t(n) = 2*n^-1` array
- for each `n` take `p = t(n)` and try to sieve it:
  - try to divide all `t(n+k*p)` by `p` as much as possible
  - try to divide all `t(-n+k*p)` by `p` as much as possible
  - consider `n` as prime generating number if current value of `p` is equal to `2*n^2-1` (meaning that this `t(n)` was not sieved meaning it is prime)

## Alternative solutions

## External links

http://devalco.de/quadr_Sieb_2x%5E2-1.php