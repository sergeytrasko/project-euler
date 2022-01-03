# Project Euler Problem 277 - A Modified Collatz sequence

[https://projecteuler.net/problem=277](https://projecteuler.net/problem=277)

## Solution idea

We can consider our answer `N` as a number in base 3 (with "digits" `D`, `U` and `d`) with slightly different exponent rules.

Start with `10^15` and try to find smallest number that matches the first symbol of a prefix, in our case it is `U`.

Then try to search next number that is greater by `3` which matches first 2 symbols of a prefix - `UD`.

Then the same process, but increase the number by `9` to match `UDD` prefix.

Then - increase by `27` and use prefix `UDDD`.

And so on...

## Alternative solutions

Set `a1 = x`. Then we can say that if sequence starts with `U` then `a2 = (4*t+2)/3`.

Next is `D` and `a3 = ((4*t+2)/3)/3`. For next `D` we get `a4 = (((4*t+2)/3)/3)/3` and so on...

What is remaining is to do some algebraic simplifications and solve for `t`.

We can notice that we can simplify the expression by noticing that denominator will be 3^30

## External links
