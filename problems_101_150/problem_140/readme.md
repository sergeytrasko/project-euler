# Project Euler Problem 140 - Modified Fibonacci golden nuggets

[https://projecteuler.net/problem=140](https://projecteuler.net/problem=140)

## Solution idea

Very similar to [problem 137 - Fibonacci golden nuggets](../problem_137) and the same problems there.

The sequence can be expressed using a generating function:
```
A(x) = (x + 3x^2) / (1 - x - x^2)
```
And if setting `A(x) = k` and doing some algebra manipulations we come to a quadratic equation:
```
x = (-(k+1) + sqrt[ (k+1)^2 + 4k(3+k) ])/(2(3+k))
```

Obviously for `x` to be a rational number `(k+1)^2 + 4k(3+k)` should be a perfect square. 
That gives us an easy brute-force algorithm:
- iterate `k` from 1 to whatever we need
- calculate `s = (k+1)^2 + 4k(3+k)`
- check if `s` is a perfect square

Similarly to problem 137 - after `k = 23` (shortly after example value) we are starting to have integer overflow problems and the calculation takes ages. 

Different approach is needed.

Let's rewrite our initial condition (`(k+1)^2 + 4k(3+k)` should be a perfect square) as follows:
- consider equation `5x^2+14x+1-y^2=0`
- `k` is replaced by `x`
- solve over positive integers as Diophantine equation

This yields 6 `(x, y)` pairs and 2 recurrent functions - in total 12 combinations that generate sequence of numbers.

I run each of them for 10 times (magic number), take only positive values of `x` and merge results together.
Next need to remove duplicates, sort the resulting array and sum first 30 members.

Having solved this problem I tried the same approach for [problem 137 - Fibonacci golden nuggets](../problem_137) and it worked nicely as well.

## Alternative solutions

## External links

https://en.wikipedia.org/wiki/Generating_function
https://www.alpertron.com.ar/QUAD.HTM