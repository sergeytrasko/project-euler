# Project Euler Problem 601 - Divisibility streaks

[https://projecteuler.net/problem=601](https://projecteuler.net/problem=601)

## Solution idea

Observations:
- `streak(2x) = 1`. Indeed: `2x` is divisible by `1`, but `2x+1` is not divisible by `2`
- seems there is no `n` such that `streak(n) = 5` and same for 9, 11, 13, 14, 17?

We can run a brute-force and notice the following pattern:
- `streak(n) = 3` => `n = 7 + 12k`
- `streak(n) = 4` => `n = 13 + 12k`
- `streak(n) = 6` => `n = 61 + 60k`
- `streak(n) = 8` => `n = 841` or `n = 1681 + 1680k` 
- `streak(n) = 10` => `n = 2521 + 2520k`
- `streak(n) = 12` => `n = 27721 + 27720k`

Factors near `k` are interesting - they look like `LCM(1, 2, 3, .., m)` where `m` is streak length.

Also from OEIS we get the following Maple algorithm that confirms our assumption:
```
N:= 1000: # to get a(1) to a(N)
A:= Vector(N, 1);
for m from 2 do
  Lm:= ilcm($1..m);
  if Lm > N then break fi;
  if Lm mod (m+1) = 0 then next fi;
  for k from 1 to floor(N/Lm) do
    A[k*Lm]:=m
  od
od:
```

Then `P(s, N) = (N-2)/lcm(1, 2, .. s)`.

However we need to pay attention that `P(5, 1000) = P(6, 1000)` - this explains why we could not find streaks of length 5 with brute force. Answer is simple - they are just longer.

So calculating `P(s, N)` we need to exclude longer streaks.

## Alternative solutions

## External links

https://oeis.org/A055874