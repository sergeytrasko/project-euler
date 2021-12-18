# Project Euler Problem 121 - Disc game prize fund

[https://projecteuler.net/problem=121](https://projecteuler.net/problem=121)

## Solution idea

Here we need to deal with probabilities carefully and apply them recursively.

First of all - how to detect the probability of winning on the last turn? Pretty simple - if number of taken blue disks is bigger than number of taken red disks, then probability of win is 100% (or 1.0), otherwise it's 0% (or 0.0).
Then imagine that at some point of time we have `red` red disks and `blue` blue disks in the bag.
The probability of taking red disk is `red/(red+blue)` and probability of taking blue disk is `blue/(red+blue)`

So we just need to sum the probabilities of taking red (and similarly blue) disks in case we proceed to the next game turn recursively.

Once we get the probability of win equal to `probability` then the result prize fund is `1/probability` (floored to the integer).

In total we need to check 2^15 cases that can be done very quickly.

The idea is very similar to [problem 151 - Paper sheets of standard sizes: an expected-value problem](../problems_151_200/problem_151)

## Alternative solutions

## External links
