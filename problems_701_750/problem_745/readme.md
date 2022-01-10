# Project Euler Problem 745 - Sum of Squares II

[https://projecteuler.net/problem=745](https://projecteuler.net/problem=745)

## Solution idea

Set `N = 10^14`.

There are `N/i^2` numbers that are divisible by `i^2`. Obviously biggest square dividing `N` (and therefore appearing in the sum) is `i^2`.
This means that the number of "square divisors" will be `sqrt(N)`.

So the solution can just check integers `i` from `1` to `sqrt(N)` and calculate how many integers are divisible by `i^2`. Most tricky part here - avoid double counting

Example: 
- `N = 100`
- Define `C(x)` as number of integers below or equal to `N` that has `x^2` as biggest factor
- Need to check at most `sqrt(N) = 10` divisors
- For simplicity start with `10`
- For `10`, `9`, and `8` we have 1 such number (`100`, `81` and `64`) - `C(10) = C(9) = C(8) = 1`
- For `7` we have 2 such numbers (`98` and `49`) - `C(7) = 2`
- For `6` we also have 2 such numbers (`72` and `36`) - `C(6) = 2`
- For `5` there are 4 such numbers (`100`, `75`, `50` and `25`). Here we see problem with `100` - it was already counted and it should be excluded
  - We need to check all numbers between `5` and `sqrt(N) = 10` if they are divisors of `5`. We get only `10` in this case and therefore should substract `C(10) = 1` from `C(5)`
  - As result `C(5) = 3`
- For `4` we have 6 such numbers (`96`, `80`, `64`, `48`, `32`, `16`)
  - Check multiples of `4` below `sqrt(N)` - it is `8` and exlcude `C(8)`
  - `C(4) = 6 - C(8) = 5`
- For `3` we have 10 such numbers (`99`, `90`, `81`, `72`, `63`, `54`, `36`, `27`, `18`, `9`)
  - Muliples of `3` are `6` and `9`
  - `C(3) = 10 - C(6) - C(9) = 10 - 2 - 1 = 7`
- For `2` we have 25 such numbers - I will not list them all, but we have problems with `16`, `36`, `72` an few others. Similarly to `5` we need to exclude few
  - Checking multiples of `2` - these are `4`, `6`, `8` and `10`
  - `C(2) = 25 - C(4) - C(6) - C(8) - C(10) = 25 - 5 - 1 - 1 = 18` 
- For `1` - same story, we have `100` numbers and we need to exclude everything above

## Alternative solutions

## External links

https://oeis.org/A008833

https://en.wikipedia.org/wiki/Inclusion%E2%80%93exclusion_principle