# Project Euler Problem 493 - Under The Rainbow

[https://projecteuler.net/problem=493](https://projecteuler.net/problem=493)

## Solution idea

My first idea was to use same approach as for [problem 121](../../problems_101_150/problem_121) or [problem 151](../../problems_151_200/problem_151). But the number of cases to check is around `7^20` that goes far beyond what modern computer can handle.

We can simplify the problem and make the following observations:
- all 7 colors are equivalent - meaning that propability for different color are the same
- if we have only 2 colors - say 10 red and 60 blue balls - then probability of choosing at least 1 red ball is `1 - C(60, 20)/C(70, 20)`. Why so? `C(70, 20)` is total number of possibilities to choose 20 balls out of 70. `C(60, 20)` - number of possibilities to choose only blue balls, these are our negative scenarios. We need the opposite - only red balls, therefore substract this fraction from one.
- Then for 7 colors the result is `7 * (1 - C(60, 20)/C(70, 20))`

Problem is that we cannot calculate `C(60, 20)` and `C(70, 20)` as they require quite big numbers to deal with as:
- `C(60, 20) = 60!/ (20! * 40!)`
- `C(70, 20) = 70!/ (20! * 50!)`

But we can simplify this expression a bit: `C(60, 20)/C(70, 20) = 60!/ (20! * 40!) / 70!/ (20! * 50!) = (41*42*...*49*50)/(61*62*...69*70)`

After that the solution becomes trivial.

## Alternative solutions

## External links

https://www.wolframalpha.com/input/?i=choose%2860%2C20%29%2Fchoose%2870%2C20%29