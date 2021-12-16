# Project Euler Problem 174 - Counting the number of "hollow" square laminae that can form one, two, three, ... distinct arrangements

[https://projecteuler.net/problem=174](https://projecteuler.net/problem=174)

## Solution idea

Actually the problem is relatively simple, but the definitions used in the problem make it quite complex to understand.

Define `L(t)` as number of ways that `t` can be used to make a lamina with any hole size.

We can fill `L(t)` by iterating `t` from 1 to 1'000'000 and trying every hole size `n` possible.

How to detect number of tiles possible? Let's take a look at a layer just around the hole. It has 4 corner tiles plus 4 times size of the hole tiles. So basically `4*n+4` where `n` is the size of the hole.
On the next row there will be again 4 corner tiles and by 2 more tiles on each horizontal and vertical sides or in other words `4*(n+2)+4`.
And so on... Basically layer `k` around the hole will contain `4*(n+(k-1)*2)+4` tiles.

And as result we need to count how many `L[i]` are there that have value between 1 and 10 inclusive.

## Alternative solutions

## External links

Quite similar to [problem 173](../problem_173)