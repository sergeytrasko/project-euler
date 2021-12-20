# Project Euler Problem 110 - Diophantine reciprocals II

[https://projecteuler.net/problem=110](https://projecteuler.net/problem=110)

## Solution idea

Despite the problem is very similar to [problem 108 - Diophantine reciprocals I](../problem_108) we need to use a more clever approach here as brute-force solution will work for years.
It can be found that the number of solutions to `1/x + 1/y = 1/n` is equal to number of divisors of `n^2`. As we are interested in only unique solutions we should divide this number by 2.
For example `1260 = 2^2 * 3^2 * 5 * 7`. Sum of divisors is `(1 + 2*2)*(1 + 2*2)*(1 + 2*1)*(1 + 2*1) = 5*5*3*3 = 225`. As we want only distinct solutions, we should ignore the symethry here meaning that the number of distinct solutions is `(225 + 1) / 2 = 113`.
Now the question boils down to: find minimal number `n^2` that has over 8'000'000-1 divisors
Instead of iterating over `1^2, 2^2, 3^2, ...` and counting number of divisors, we can try to construct such number `n^2`.
Each positive integer can be written as `p1^a1 * p2^a2 * ... pk^ak` where `pi` are distinct primes. We need to choose primes `p` and their powers `a` so that:
- the product is as small as possible
- `mul[ (1 + 2ai) ] > 4'000'000`

First of all: how many primes do we need at most? It turns out that if we use less than 14 primes then `(1+2*a1)*...(1+2*a14) < 8'000'000`. This makes us think that `n` should consist of not more than 15 prime factors.
Then - how to arange the powers? I make the following claim: if `pi < pj` then `ai >= aj`. That's pretty obvious as the placement of powers does not influence number of divisors, but it influence the value of `n`.

It turns out that we can do it recursively.
- We know that we need at most 15 primes
- Start with smallest prime and smallest power of it - `2^1`
- On each recursive step there are 2 options what to do next:
  - increase the power (e.g. from `2^5` to `2^6`)
  - append another prime to the list of options
- Keep power rule in place: if `pi < pj` then `ai >= aj`
- Keep track of "best" option found so far and cut off the results that are already worse

Note that this solution can be used to solve [problem 108 - Diophantine reciprocals I](../problem_108) as well.

## Alternative solutions

## External links

https://mathforums.com/threads/diophantine-reciprocals-amount-of-solutions.42193/
